import java.io.DataInputStream;
import java.io.FileInputStream;
import java.nio.file.Files;
import java.nio.file.Paths;

/**
 * Independent JDK 11 -Xverify:all oracle for the original Fliplay class files.
 * Input is an already hash-checked original .flc file in a temporary folder.
 * The historical bytecode parses the header, loads every frame, and decodes
 * each frame. This harness only supplies the headless state and exports bytes;
 * it does not run the original Applet UI, thread, AWT renderer or file URL.
 */
public final class FliplayAllFramesJdk {
  public static void main(String[] args) throws Exception {
    if (args.length != 2) throw new IllegalArgumentException("FLC and output directory required");
    fliplay player = new fliplay();
    player.red = new byte[256];
    player.green = new byte[256];
    player.blue = new byte[256];
    try (DataInputStream in = new DataInputStream(new FileInputStream(args[0]))) {
      player.load_header(in);
      if (player.fli_frames != 60 || player.image_width != 320 || player.image_height != 240) {
        throw new IllegalStateException("original header disagrees with pinned input");
      }
      player.pixels = new byte[player.image_width * player.image_height];
      player.row_start = new int[player.image_height];
      player.row_end = new int[player.image_height];
      player.frame = new flickframe[player.fli_frames + 1];
      for (int i = 0; i <= player.fli_frames; ++i) {
        // Original flickframe constructor supplies AWT patch arrays, unused by
        // load_frame/extract_frame; the decoder reads only the chunks field.
        player.frame[i] = new flickframe(1);
        player.load_frame(in, i);
        player.extract_frame(i);
        String id = String.format("%03d", i);
        Files.write(Paths.get(args[1], id + ".idx"), player.pixels);
        byte[] palette = new byte[768];
        System.arraycopy(player.red, 0, palette, 0, 256);
        System.arraycopy(player.green, 0, palette, 256, 256);
        System.arraycopy(player.blue, 0, palette, 512, 256);
        Files.write(Paths.get(args[1], id + ".palette"), palette);
        System.out.println("JDK original load_frame/extract_frame " + i);
      }
      if (in.read() != -1) throw new IllegalStateException("original loader left trailing bytes");
    }
  }
}
