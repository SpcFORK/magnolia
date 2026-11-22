"""Owned finite FWS1 semantic parser tests; real SDK C++ oracle is separate."""
from __future__ import annotations

import importlib.util
import os
from pathlib import Path
import struct
import subprocess
import sys
import tempfile
import unittest

ROOT=Path(__file__).resolve().parents[1]
SPEC=importlib.util.spec_from_file_location('fws1_semantic_dump',
                                             ROOT/'tools/swf-sdk-dump.py')
assert SPEC and SPEC.loader
sdk=importlib.util.module_from_spec(SPEC)
sys.modules[SPEC.name]=sdk
SPEC.loader.exec_module(sdk)


class BitWriter:
    def __init__(self):self.bits=[]
    def put(self,value:int,width:int):
        self.bits.extend((value>>at)&1 for at in range(width-1,-1,-1))
        return self
    def bytes(self):
        self.bits.extend([0]*(-len(self.bits)%8))
        return bytes(sum(self.bits[i+j]<<(7-j) for j in range(8))
                     for i in range(0,len(self.bits),8))


def rect():
    return (BitWriter().put(9,5).put(0,9).put(200,9)
            .put(0,9).put(200,9).bytes())


def swf_tag(code:int,payload:bytes=b'')->bytes:
    if len(payload)>=63:raise AssertionError('test tags must fit short header')
    return struct.pack('<H',(code<<6)|len(payload))+payload


def owned_movie(*,curve:bool=False,shape_suffix:bytes=b'')->bytes:
    path=(BitWriter().put(1,4).put(0,4)  # fill width=1, line width=0
          .put(0,1).put(0b00101,5)       # move + FillStyle1
          .put(1,5).put(0,1).put(0,1).put(1,1))
    if curve:
        path=(path.put(1,1).put(0,1).put(3,4)  # curved edge, 5-bit deltas
              .put(5,5).put(10,5).put(5,5).put(-10&31,5))
    else:
        path=(path.put(1,1).put(1,1).put(6,4)  # straight 8-bit horizontal
              .put(0,1).put(0,1).put(100,8))
    path=path.put(0,1).put(0,5).bytes()  # EndShape
    shape=(struct.pack('<H',1)+rect()+b'\x01\x00\xff\x00\x00\x00'+path+
           shape_suffix)
    tags=swf_tag(2,shape)+swf_tag(1)+swf_tag(0)
    movie=b'FWS\x01'+b'\0'*4+rect()+struct.pack('<HH',12*256,1)+tags
    return movie[:4]+struct.pack('<I',len(movie))+movie[8:]


class FiniteSWF1SemanticTest(unittest.TestCase):
    def test_owned_real_shape_line_and_frame(self):
        result=sdk.parse(owned_movie())
        self.assertEqual((result['format'],result['frames'],len(result['tags'])),
                         ('FWS1',1,3))
        shape=result['shapes'][1]
        self.assertEqual(shape['styleGroups'][0]['solidFillsRGB'],['ff0000'])
        self.assertEqual(shape['events'],[
            {'op':'styles','moveTwips':[0,0],'fill1':1},
            {'op':'horizontal','endTwips':[100,0]}])

    def test_owned_curved_edge_has_independent_control_and_anchor(self):
        shape=sdk.parse(owned_movie(curve=True))['shapes'][1]
        self.assertEqual(shape['events'][-1],
                         {'op':'curve','controlTwips':[5,10],'endTwips':[10,0]})

    def test_malformed_payload_unsupported_tag_and_pin_refuse_without_artifact(self):
        invalid={
            'undeclared shape tail':owned_movie(shape_suffix=b'\x01'),
            'unsupported control opcode':owned_movie()[:-2]+b'\x80\x00',
            'bad declared length':owned_movie()[:4]+b'\xff'*4+owned_movie()[8:],
        }
        with tempfile.TemporaryDirectory() as temp:
            for name,data in invalid.items():
                inp=Path(temp)/(name+'.swf');out=Path(temp)/(name+'.json')
                inp.write_bytes(data)
                proc=subprocess.run([sys.executable,str(ROOT/'tools/swf-sdk-dump.py'),
                    '--input',str(inp),'--output',str(out)],
                    capture_output=True,text=True,timeout=5,cwd=ROOT)
                self.assertEqual(proc.returncode,2,msg=name+proc.stderr)
                self.assertIn('FWS1 parser refused before output:',proc.stderr)
                self.assertFalse(out.exists())
            inp=Path(temp)/'pin.swf';out=Path(temp)/'pin.json'
            inp.write_bytes(owned_movie())
            proc=subprocess.run([sys.executable,str(ROOT/'tools/swf-sdk-dump.py'),
                '--input',str(inp),'--output',str(out),'--expect-sha256','0'*64],
                capture_output=True,text=True,timeout=5,cwd=ROOT)
            self.assertEqual(proc.returncode,2)
            self.assertFalse(out.exists())


if __name__=='__main__':unittest.main()
