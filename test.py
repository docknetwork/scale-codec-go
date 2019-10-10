from ctypes import *

from scalecodec.base import ScaleBytes, ScaleDecoder
from scalecodec import U32

lib = cdll.LoadLibrary("./cdecoder.so")


# class GoInput(Structure):
#     _fields_ = [("cv", c_char_p)]

class GoRes(Structure):
    _fields_ = [("ci", c_uint), ("cerr", c_char_p)]


lib.CDecodeUInt32.argtypes = [c_char_p]
lib.CDecodeUInt32.restype = GoRes

# var = "0x02093d00"
# print(type(var.encode("utf-8")))
from contextlib import contextmanager


@contextmanager
def time_it():
    from time import time
    t = time()
    try:
        yield
    finally:
        print("spent", time() - t)


v = "0x02093d00"
# v = "0xfeffffff"


obj = ScaleDecoder.get_decoder_class('Compact<u32>', ScaleBytes(v))
with time_it():
    for _ in range(10000):
        obj = U32(ScaleBytes(v))
        obj = ScaleDecoder.get_decoder_class('Compact<u32>', ScaleBytes(v))
        obj.decode()
        obj.value

with time_it():
    for _ in range(10000):
        res = lib.CDecodeUInt32(v.encode('utf-8'))
        res.ci

