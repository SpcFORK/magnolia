#!/usr/bin/env python3
"""Replay all 60 Moon Clouds frames through ORIGINAL Fliplay JVM bytecode.

The user's SHA-256-pinned FLC archive and the unmodified archived class files
stay external to Git. The native Oak VM executes original load_header,
read_loop, load_frame, extract_frame, BRUN/palette/DELTA_FLC decoder methods;
the runner supplies bounded in-memory streams, minimal frame scaffolding and
an export sink. A separate JDK 11 -Xverify:all process invokes the same original
methods and an independent Pillow FLC decoder checks EVERY indexed/RGB frame.
The optional offline HTML displays lossless exports, not historical Applet UI.
"""
from __future__ import annotations

import argparse
import base64
import hashlib
from io import BytesIO
import json
from pathlib import Path
import re
import subprocess
import sys
import tempfile
import time
from zipfile import ZipFile

from PIL import Image

ROOT = Path(__file__).resolve().parent.parent
PINNED = {
    "zip": "15b5d795fd96ae4cf32be8610745fff21611407803b2d3e59b3a430f4971c564",
    "flc": "812379899d151a1460366249fd794899eeefe3c7bacdc528da22c2624b33cff4",
    "fliplay": "bc315872dafebbf3c634e64c0fc251b38031ed49b7c4698d39c94990de9fe20b",
    "flickframe": "12feb5321cd8c9aeb0bf285f055596d981e3d45ccfb8c62d029f2f6d95ccdf9f",
    "link": "3fce7191fa4d0d657b33d017647bdc96e160f77cba061a7c90e60e88e7138fa5",
    "indices-all": "dd06653c96de9da8cf3b066c72a9885ef3c19d5a8c5de662b42cb4b6a4c26102",
    "rgb-all": "dfb9f2fbde80dafbee976ae43ec2c04c1019d7ffccb13807400324465f89fe16",
}
HEADER_TEXT = "File Size:  366503\nMagic:      44818\nFrames:     60\n320x240\n"
FRAME_LINE = re.compile(
    r"^\[fliplay-all-frames\] FRAME (\d{3}) bytes=(\d+) chunks=(\d+) "
    r"read=(\d+) steps=(\d+) cells=(\d+)$", re.M
)
COMPLETE_LINE = re.compile(
    r"^\[fliplay-all-frames\] COMPLETE 61 loaded/extracted frames; "
    r"steps=(\d+) cells=(\d+) \(this is not original AWT playback\)$", re.M
)


def sha(data: bytes) -> str:
    return hashlib.sha256(data).hexdigest()


def checked(path: Path, kind: str, max_bytes: int) -> bytes:
    if not path.is_file() or path.stat().st_size > max_bytes:
        raise ValueError(f"missing/oversized {kind}: {path}")
    raw = path.read_bytes()
    if sha(raw) != PINNED[kind]:
        raise ValueError(f"{kind} SHA-256 mismatch: {sha(raw)}")
    return raw


def u16(raw: bytes, at: int) -> int:
    return int.from_bytes(raw[at:at + 2], "little")


def u32(raw: bytes, at: int) -> int:
    return int.from_bytes(raw[at:at + 4], "little")


def validate_flc(raw: bytes) -> list[dict]:
    if len(raw) != 366503 or sha(raw) != PINNED["flc"]:
        raise ValueError("original FLC size/digest mismatch")
    if (u32(raw, 0), u16(raw, 4), u16(raw, 6), u16(raw, 8),
            u16(raw, 10), u16(raw, 12), u32(raw, 16)) != (
                366503, 0xAF12, 60, 320, 240, 8, 71):
        raise ValueError("original FLC header geometry/timing mismatch")
    offset, frames = 128, []
    for i in range(61):  # fliplay.load_fli includes the trailing 16-byte ring
        if offset + 16 > len(raw):
            raise ValueError(f"short frame {i} header")
        size, magic, count = u32(raw, offset), u16(raw, offset + 4), u16(raw, offset + 6)
        if magic != 0xF1FA or not 16 <= size <= 16384 or count > 2 or \
                offset + size > len(raw):
            raise ValueError(f"invalid bounded frame {i}")
        at, kinds = offset + 16, []
        for _ in range(count):
            if at + 6 > offset + size:
                raise ValueError(f"short chunk header in frame {i}")
            chunk_size, kind = u32(raw, at), u16(raw, at + 4)
            if not 6 <= chunk_size <= 14000 or at + chunk_size > offset + size:
                raise ValueError(f"invalid chunk in frame {i}")
            kinds.append(kind)
            at += chunk_size
        if at != offset + size:
            raise ValueError(f"unparsed bytes in frame {i}")
        expected = [11, 15] if i == 0 else ([] if i == 60 else [7])
        if kinds != expected or (i == 60 and size != 16):
            raise ValueError(f"unexpected chunk types in frame {i}: {kinds}")
        frames.append({"number": i, "offset": offset, "bytes": size,
                       "chunks": count, "chunkTypes": kinds})
        offset += size
    if offset != len(raw):
        raise ValueError("unparsed FLC tail")
    return frames


def movie(path: Path) -> tuple[bytes, list[dict]]:
    checked(path, "zip", 1_000_000)
    with ZipFile(path) as z:
        infos = z.infolist()
        if len(infos) != 1 or infos[0].filename != "moonclouds.flc" or \
                infos[0].file_size != 366503 or infos[0].flag_bits & 1:
            raise ValueError("expected exactly one unencrypted bounded moonclouds.flc")
        raw = z.read(infos[0])
    return raw, validate_flc(raw)


def run(label: str, argv: list[str], cwd: Path, timeout: int) -> str:
    print(f"[fliplay-all-frames] START {label} (cap {timeout}s)", flush=True)
    t0 = time.monotonic()
    process = subprocess.run(argv, cwd=cwd, stdout=subprocess.PIPE,
                             stderr=subprocess.STDOUT, text=True, timeout=timeout)
    if process.returncode:
        raise RuntimeError(f"{label} exited {process.returncode}:\n{process.stdout[-2500:]}")
    print(f"[fliplay-all-frames] PASS {label} ({time.monotonic() - t0:.1f}s)", flush=True)
    return process.stdout


def check_logs(jdk: str, oak: str, frames: list[dict]) -> dict:
    expected_jdk = HEADER_TEXT + "".join(
        f"JDK original load_frame/extract_frame {i}\n" for i in range(61))
    if jdk != expected_jdk:
        raise ValueError(f"JDK original-method log diverged: {jdk[:500]!r}")
    lines = [tuple(map(int, m.groups())) for m in FRAME_LINE.finditer(oak)]
    done = COMPLETE_LINE.search(oak)
    if len(lines) != 61 or done is None or oak.count("[fliplay-all-frames] COMPLETE") != 1:
        raise ValueError(f"Oak did not complete all 61 frame/ring calls: {oak[-2000:]}")
    prior_steps = prior_cells = 0
    for (number, size, chunks, read, steps, cells), frame in zip(lines, frames):
        if (number, size, chunks, read) != (
                frame["number"], frame["bytes"], frame["chunks"], frame["bytes"]):
            raise ValueError(f"Oak frame {number} length/chunks/read mismatch")
        if not prior_steps < steps <= 10_000_000 or \
                not prior_cells <= cells <= 1_048_576:
            raise ValueError(f"Oak frame {number} budget non-monotonic/violated")
        prior_steps, prior_cells = steps, cells
    if tuple(map(int, done.groups())) != (prior_steps, prior_cells):
        raise ValueError("Oak cumulative budgets disagree with final line")
    return {"instructions": prior_steps, "chargedArrayCells": prior_cells,
            "perArrayCap": 81920, "stepCap": 10_000_000,
            "cumulativeArrayCellCap": 1_048_576,
            "loadedFrameHeaders": 61, "displayedFrames": 60}


def verify_frames(jdk_dir: Path, oak_dir: Path, raw: bytes,
                  template: str | None) -> tuple[list[dict], bytes | None, dict]:
    with Image.open(BytesIO(raw)) as independent:
        if (independent.mode, independent.size, independent.n_frames,
                independent.info.get("duration")) != ("P", (320, 240), 60, 71):
            raise ValueError("independent Pillow FLC decoder metadata disagrees")
        data_urls, frame_hashes, report = [], [], []
        idx_aggregate, rgb_aggregate = hashlib.sha256(), hashlib.sha256()
        prior_pixels = None
        last_pixels = last_palette = None
        for i in range(61):
            name = f"{i:03d}"
            indexed = (oak_dir / f"{name}.idx").read_bytes()
            palette = (oak_dir / f"{name}.palette").read_bytes()
            if len(indexed) != 76800 or len(palette) != 768:
                raise ValueError(f"Oak frame {i} has wrong array size")
            if indexed != (jdk_dir / f"{name}.idx").read_bytes() or \
                    palette != (jdk_dir / f"{name}.palette").read_bytes():
                raise ValueError(f"Oak/JDK original-bytecode pixels or palette differ at frame {i}")
            if i == 60:
                if indexed != last_pixels or palette != last_palette:
                    raise ValueError("empty ring frame changed output")
                continue
            independent.seek(i)
            rgb = bytes(channel for pixel in indexed for channel in (
                palette[pixel], palette[256 + pixel], palette[512 + pixel]))
            if indexed != independent.tobytes() or rgb != independent.convert("RGB").tobytes():
                raise ValueError(f"original decoded frame {i} differs from Pillow FLC")
            idx_aggregate.update(indexed)
            rgb_aggregate.update(rgb)
            changed = 0 if prior_pixels is None else sum(
                a != b for a, b in zip(indexed, prior_pixels))
            item = {"frame": i, "indexedSHA256": sha(indexed),
                    "paletteSHA256": sha(palette), "rgbSHA256": sha(rgb),
                    "changedIndexedPixelsFromPrevious": changed}
            if template is not None:
                # Pillow is used here solely as a PNG *encoder*. Its FLC
                # decoder is the independent comparison above; the source
                # pixels/palette in the HTML come from Magnolia's original VM.
                interleaved_palette = bytes(
                    value for n in range(256) for value in (
                        palette[n], palette[n + 256], palette[n + 512]))
                png = Image.frombytes("P", (320, 240), indexed)
                png.putpalette(interleaved_palette)
                output = BytesIO()
                png.save(output, format="PNG", optimize=True)
                packed = output.getvalue()
                with Image.open(BytesIO(packed)) as decoded:
                    if decoded.convert("RGB").tobytes() != rgb:
                        raise ValueError(f"lossless PNG roundtrip failed at {i}")
                item["embeddedPngSHA256"] = sha(packed)
                data_urls.append("data:image/png;base64," + base64.b64encode(packed).decode("ascii"))
                frame_hashes.append(item["indexedSHA256"])
            report.append(item)
            prior_pixels, last_pixels, last_palette = indexed, indexed, palette
    if idx_aggregate.hexdigest() != PINNED["indices-all"] or \
            rgb_aggregate.hexdigest() != PINNED["rgb-all"]:
        raise ValueError("aggregate all-frame SHA-256 regression")
    summary = {"allIndexedSHA256": idx_aggregate.hexdigest(),
               "allRGBSHA256": rgb_aggregate.hexdigest(),
               "changedIndexedPixelsAcrossLoop": sum(x["changedIndexedPixelsFromPrevious"] for x in report),
               "ringPreservesFinalPixelsAndPalette": True}
    html = None
    if template is not None:
        if template.count("__FRAME_DATA__") != 1 or \
                template.count("__FIRST_FRAME__") != 1 or \
                template.count("__FIRST_HASH__") != 1:
            raise ValueError("HTML template placeholders changed")
        content = template.replace("__FRAME_DATA__", json.dumps(
            {"frames": data_urls, "hashes": frame_hashes}, separators=(",", ":")))
        content = content.replace("__FIRST_FRAME__", data_urls[0])
        content = content.replace("__FIRST_HASH__", frame_hashes[0])
        html = content.encode("utf-8")
        summary["offlineHtmlSHA256"] = sha(html)
        summary["offlineHtmlBytes"] = len(html)
    return report, html, summary


def main() -> None:
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("--movie", type=Path, required=True, help="uploaded moonclouds.zip.txt")
    parser.add_argument("--fliplay", type=Path, required=True, help="hash-pinned original fliplay.class")
    parser.add_argument("--flickframe", type=Path, required=True, help="hash-pinned original flickframe.class")
    parser.add_argument("--magnolia", type=Path, required=True, help="built Magnolia interpreter")
    parser.add_argument("--html", type=Path, help="offline viewer from Oak's decoded pixels (optional)")
    parser.add_argument("--report", type=Path, help="per-frame SHA-256 evidence (optional)")
    args = parser.parse_args()
    java_version = subprocess.run(["java", "-version"], capture_output=True,
                                  text=True, check=True).stderr
    javac_version = subprocess.run(["javac", "-version"], capture_output=True,
                                   text=True, check=True).stdout
    if 'version "11' not in java_version or not javac_version.startswith("javac 11"):
        raise ValueError("JDK 11 java/javac required for this differential oracle")
    if not args.magnolia.is_file():
        raise ValueError(f"missing Magnolia executable: {args.magnolia}")
    raw, frames = movie(args.movie)
    originals = {kind: checked(getattr(args, kind), kind, 100_000)
                 for kind in ("fliplay", "flickframe")}
    link = ROOT / "test/historical/FliplayFrameLink.class"
    checked(link, "link", 100_000)
    template = (ROOT / "tools/fliplay-original-viewer-template.html").read_text() \
        if args.html is not None else None
    with tempfile.TemporaryDirectory(prefix="magnolia-original-fliplay-all-") as scratch:
        work = Path(scratch)
        originals_dir, input_dir = work / "classes", work / "frames"
        jdk_dir, oak_dir = work / "jdk", work / "oak"
        for directory in (originals_dir, input_dir, jdk_dir, oak_dir, work / "sourcepath"):
            directory.mkdir()
        for kind, bytes_ in originals.items():
            (originals_dir / f"{kind}.class").write_bytes(bytes_)
        (input_dir / "original.flc").write_bytes(raw)
        (input_dir / "header.bin").write_bytes(raw[:128])
        for frame in frames:
            at = frame["offset"]
            (input_dir / f'{frame["number"]:03d}.bin').write_bytes(
                raw[at:at + frame["bytes"]])
        run("JDK 11 javac --release 8 builds independent harness and reproducible link",
            ["javac", "-proc:none", "--release", "8", "-g:none", "-sourcepath",
             str(work / "sourcepath"), "-cp", str(originals_dir), "-d",
             str(originals_dir), str(ROOT / "test/historical/FliplayFrameLink.java"),
             str(ROOT / "test/historical/FliplayAllFramesJdk.java")], ROOT, 40)
        if (originals_dir / "FliplayFrameLink.class").read_bytes() != link.read_bytes():
            raise ValueError("Java 8 reachability link did not rebuild byte-for-byte")
        jdk = run("JDK 11 -Xverify:all original load_header/load_frame/extract_frame 60 + ring",
            ["java", "-Xverify:all", "-Djava.awt.headless=false", "-cp",
             str(originals_dir), "FliplayAllFramesJdk",
             str(input_dir / "original.flc"), str(jdk_dir)], ROOT, 80)
        oak = run("Magnolia Oak original bytecode 60 frames + ring",
            [str(args.magnolia), str(ROOT / "tools/fliplay-all-frames.oak"),
             "--frames", str(input_dir), "--out", str(oak_dir),
             "--fliplay", str(originals_dir / "fliplay.class"),
             "--flickframe", str(originals_dir / "flickframe.class"),
             "--link", str(link)], ROOT, 900)
        budgets = check_logs(jdk, oak, frames)
        per_frame, html, summary = verify_frames(jdk_dir, oak_dir, raw, template)
        print("[fliplay-all-frames] PASS all 60 original decoded indexed frames + palettes "
              "match JDK original methods; all 60 RGB frames match independent Pillow; "
              "ring frame preserved", flush=True)
    report = {
        "sourceArchiveSHA256": PINNED["zip"], "sourceFLC_SHA256": PINNED["flc"],
        "originalClassSHA256": {key: PINNED[key] for key in ("fliplay", "flickframe")},
        "reachabilityFixtureSHA256": PINNED["link"],
        "format": {"width": 320, "height": 240, "indexedBitsPerPixel": 8,
                   "frames": 60, "frameDelayMs": 71, "ringHeaderBytes": 16},
        "execution": {"historicalMethodsActuallyUsed": [
            "fliplay.<init>", "fliplay.load_header", "fliplay.read_loop",
            "fliplay.get_ulong", "fliplay.get_ushort", "fliplay.get_ubyte",
            "fliplay.load_frame", "fliplay.extract_frame",
            "fliplay.reset_row_start_end", "fliplay.process_color_256_chunk",
            "fliplay.process_brun_chunk", "fliplay.process_delta_chunk"],
            "linkedButNotTriggeredForThisMovie": ["fliplay.process_lc_chunk"],
            "streams": "Python validates/splits uploaded FLC; JVM VM uses bounded in-memory byte streams per header/frame",
            "frameObjects": "VM supplies default flickframe objects; original load_frame writes chunks; AWT patch constructor is not run",
            "viewer": "new HTML playback shell with precomputed Oak decoder frames, not original AWT/Java browser execution",
            "originalAppletFileOpeningOrThreadOrAWTRendering": False,
            "flashSWFOrProjectorPlayback": False, "budgets": budgets},
        "verification": {"jdk": "JDK 11 -Xverify:all original Fliplay methods on continuous .flc stream",
                         "independentDecoder": "Pillow FLC: 60 indexed frames and RGB bytes pixel-for-pixel",
                         "oakVsJdkIndexedAndPalette": "61/61 including empty ring",
                         "oakVsPillowIndexedAndRGB": "60/60", **summary},
        "frames": per_frame,
    }
    if args.html is not None and html is not None:
        args.html.parent.mkdir(parents=True, exist_ok=True)
        args.html.write_bytes(html)
        print(f"[fliplay-all-frames] wrote offline original-decoder frame viewer: {args.html} "
              f"({len(html):,} bytes)", flush=True)
    if args.report is not None:
        args.report.parent.mkdir(parents=True, exist_ok=True)
        args.report.write_text(json.dumps(report, indent=2) + "\n")
        print(f"[fliplay-all-frames] wrote evidence report: {args.report}", flush=True)
    print("[fliplay-all-frames] DONE original load/decode complete for this movie; "
          "historical Applet UI, URL/file opener, AWT image rendering, animation "
          "thread, Flash/SWF and native projector remain unimplemented", flush=True)


if __name__ == "__main__":
    try:
        main()
    except (ValueError, RuntimeError, OSError, subprocess.TimeoutExpired) as error:
        print(f"[fliplay-all-frames] FAIL: {error}", file=sys.stderr)
        sys.exit(1)
