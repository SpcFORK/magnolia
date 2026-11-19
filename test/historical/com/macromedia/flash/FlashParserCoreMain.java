package com.macromedia.flash;

// JDK oracle for three methods of the *original* FlashApi.ZIP SParser class.
// Their obfuscated names, here expressed as Java Unicode escapes, perform
// little-endian unsigned/signed reads from the original bytecode. The fields
// and methods intentionally share names: that is how the old class is stored.
public final class FlashParserCoreMain {
    public static void main(String[] args) {
        SParser parser = new SParser();
        parser.\u00c1 = new byte[] {(byte) 0xff, (byte) 0x80,
            (byte) 0xff, (byte) 0xff, 0x34, 0x12, 0, (byte) 0x80};
        parser.\u00c2 = 0;
        System.out.println(parser.\u00c1());
        System.out.println(parser.\u00c2());
        System.out.println(parser.\u00c3());
        System.out.println(parser.\u00c2);
    }
}
