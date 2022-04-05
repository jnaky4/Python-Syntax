class B_Type_Literals {
    static void main(String[] args) {
        /*
            Number Type Suffixes
            We can force a number (including binary, octals and hexadecimals) to have a specific type by giving a suffix
            (see print below), either uppercase or lowercase.
         */
        def type_suffix = """\
        Integer     I or i  ${42I} ${42i} ${"Hex->Int: " + 0xFFi}
        Long        L or l  ${42L} ${"Byte->Long: " + 0b1111L}
        BigInteger  G or g  ${456G} ${"Octal->Int: " + 034G}
        BigDecimal  G or g  ${456g}
        Double      D or d  ${1.23E23D}
        Float       F or f  ${123.4F}
"""
        println type_suffix

        //Underscore Literals
        long creditCardNumber = 1234_5678_9012_3456L
        long socialSecurityNumbers = 999_99_9999L
        double monetaryAmount = 12_345_132.12
        long hexBytes = 0xFF_EC_DE_5E
        long hexWords = 0xFFEC_DE5E
        long maxLong = 0x7fff_ffff_ffff_ffffL
        long alsoMaxLong = 9_223_372_036_854_775_807L
        long bytes = 0b11010010_01101001_10010100_10010010


        /*
            Unlike Java, Groovy doesn’t have an explicit character literal.
            However, you can be explicit about making a Groovy string an actual character, by three different means:
         */
        char c1 = 'A'
        assert c1 instanceof Character

        def c2 = 'B' as char
        assert c2 instanceof Character

        def c3 = (char)'C'
        assert c3 instanceof Character

        //Integral Literals
        // primitive types
        byte  b = 1
        char  c = 2
        short s = 3
        int   i = 4
        long  l = 5

        // infinite precision
        BigInteger bi =  6

        //Binary Literal
        int xInt = 0b10101111
        assert xInt == 175

        short xShort = 0b11001001
        assert xShort == 201 as short

        byte xByte = 0b11
        assert xByte == 3 as byte

        long xLong = 0b101101101101
        assert xLong == 2925l

        BigInteger xBigInteger = 0b111100100001
        assert xBigInteger == 3873g

        int xNegativeInt = -0b10101111
        assert xNegativeInt == -175

        //Octal literal
        //Octal numbers are specified in the typical format of 0 followed by octal digits.
        int octxInt = 077
        assert octxInt == 63

        short octxShort = 011
        assert octxShort == 9 as short

        byte octxByte = 032
        assert octxByte == 26 as byte

        long octxLong = 0246
        assert octxLong == 166l

        BigInteger octxBigInteger = 01111
        assert octxBigInteger == 585g

        int octxNegativeInt = -077
        assert octxNegativeInt == -63

        //Hexadecimal literal
        int hexxInt = 0x77
        assert hexxInt == 119

        short hexxShort = 0xaa
        assert hexxShort == 170 as short

        byte hexxByte = 0x3a
        assert hexxByte == 58 as byte

        long hexxLong = 0xffff
        assert hexxLong == 65535l

        BigInteger hexxBigInteger = 0xaaaa
        assert hexxBigInteger == 43690g

        Double hexxDouble = new Double('0x1.0p0')
        assert hexxDouble == 1.0d

        int hexxNegativeInt = -0x77
        assert hexxNegativeInt == -119

    }
}
