package utr50

var table = []struct {
	from rune
	to   rune
	prop Prop
}{
	{0x0, 0x1F, R},
	{0x20, 0x20, R},
	{0x21, 0x21, R},
	{0x22, 0x22, R},
	{0x23, 0x23, R},
	{0x24, 0x24, R},
	{0x25, 0x25, R},
	{0x26, 0x26, R},
	{0x27, 0x27, R},
	{0x28, 0x28, R},
	{0x29, 0x29, R},
	{0x2A, 0x2A, R},
	{0x2B, 0x2B, R},
	{0x2C, 0x2C, R},
	{0x2D, 0x2D, R},
	{0x2E, 0x5A, R},
	{0x5B, 0x5B, R},
	{0x5C, 0x5C, R},
	{0x5D, 0x5D, R},
	{0x5E, 0x5E, R},
	{0x5F, 0x5F, R},
	{0x60, 0x7A, R},
	{0x7B, 0x7B, R},
	{0x7C, 0x7C, R},
	{0x7D, 0x7D, R},
	{0x7E, 0x7E, R},
	{0x7F, 0x7F, R},
	{0x80, 0x9F, R},
	{0xA0, 0xA0, R},
	{0xA1, 0xA1, R},
	{0xA2, 0xA2, R},
	{0xA3, 0xA3, R},
	{0xA4, 0xA4, R},
	{0xA5, 0xA5, R},
	{0xA6, 0xA6, R},
	{0xA7, 0xA7, U},
	{0xA8, 0xA8, R},
	{0xA9, 0xA9, U},
	{0xAA, 0xAA, R},
	{0xAB, 0xAB, R},
	{0xAC, 0xAC, R},
	{0xAD, 0xAD, R},
	{0xAE, 0xAE, U},
	{0xAF, 0xAF, R},
	{0xB0, 0xB0, R},
	{0xB1, 0xB1, U},
	{0xB2, 0xB2, R},
	{0xB3, 0xB3, R},
	{0xB4, 0xB4, R},
	{0xB5, 0xB5, R},
	{0xB6, 0xB6, R},
	{0xB7, 0xB7, R},
	{0xB8, 0xB8, R},
	{0xB9, 0xB9, R},
	{0xBA, 0xBA, R},
	{0xBB, 0xBB, R},
	{0xBC, 0xBC, U},
	{0xBD, 0xBD, U},
	{0xBE, 0xBE, U},
	{0xBF, 0xBF, R},
	{0xC0, 0xD6, R},
	{0xD7, 0xD7, U},
	{0xD8, 0xF6, R},
	{0xF7, 0xF7, U},
	{0xF8, 0xFF, R},
	{0x100, 0x17F, R},
	{0x180, 0x24F, R},
	{0x250, 0x2AF, R},
	{0x2B0, 0x2E4, R},
	{0x2E5, 0x2E5, R},
	{0x2E6, 0x2E6, R},
	{0x2E7, 0x2E7, R},
	{0x2E8, 0x2E8, R},
	{0x2E9, 0x2E9, R},
	{0x2EA, 0x2EA, U},
	{0x2EB, 0x2EB, U},
	{0x2EC, 0x2FF, R},
	{0x300, 0x36F, R},
	{0x370, 0x3FF, R},
	{0x400, 0x4FF, R},
	{0x500, 0x52F, R},
	{0x530, 0x589, R},
	{0x58A, 0x58A, R},
	{0x58B, 0x58C, R},
	{0x58D, 0x58E, R},
	{0x58F, 0x58F, R},
	{0x590, 0x5BD, R},
	{0x5BE, 0x5BE, R},
	{0x5BF, 0x5FF, R},
	{0x600, 0x6FF, R},
	{0x700, 0x74F, R},
	{0x750, 0x77F, R},
	{0x780, 0x7BF, R},
	{0x7C0, 0x7FF, R},
	{0x800, 0x83F, R},
	{0x840, 0x85F, R},
	{0x8A0, 0x8FF, R},
	{0x900, 0x97F, R},
	{0x980, 0x9FF, R},
	{0xA00, 0xA7F, R},
	{0xA80, 0xAFF, R},
	{0xB00, 0xB7F, R},
	{0xB80, 0xBFF, R},
	{0xC00, 0xC7F, R},
	{0xC80, 0xCFF, R},
	{0xD00, 0xD7F, R},
	{0xD80, 0xDFF, R},
	{0xE00, 0xE7F, R},
	{0xE80, 0xEFF, R},
	{0xF00, 0xFFF, R},
	{0x1000, 0x109F, R},
	{0x10A0, 0x10FF, R},
	{0x1100, 0x11FF, U},
	{0x1200, 0x137F, R},
	{0x1380, 0x139F, R},
	{0x13A0, 0x13FF, R},
	{0x1400, 0x1400, R},
	{0x1401, 0x167F, U},
	{0x1680, 0x169F, R},
	{0x16A0, 0x16FF, R},
	{0x1700, 0x171F, R},
	{0x1720, 0x173F, R},
	{0x1740, 0x175F, R},
	{0x1760, 0x177F, R},
	{0x1780, 0x17FF, R},
	{0x1800, 0x18AF, R},
	{0x18B0, 0x18FF, U},
	{0x1900, 0x194F, R},
	{0x1950, 0x197F, R},
	{0x1980, 0x19DF, R},
	{0x19E0, 0x19FF, R},
	{0x1A00, 0x1A1F, R},
	{0x1A20, 0x1AAF, R},
	{0x1AB0, 0x1AFF, R},
	{0x1B00, 0x1B7F, R},
	{0x1B80, 0x1BBF, R},
	{0x1BC0, 0x1BFF, R},
	{0x1C00, 0x1C4F, R},
	{0x1C50, 0x1C7F, R},
	{0x1C80, 0x1C8F, R},
	{0x1CC0, 0x1CCF, R},
	{0x1CD0, 0x1CFF, R},
	{0x1D00, 0x1D7F, R},
	{0x1D80, 0x1DBF, R},
	{0x1DC0, 0x1DFF, R},
	{0x1E00, 0x1EFF, R},
	{0x1F00, 0x1FFF, R},
	{0x2000, 0x200A, R},
	{0x200B, 0x200F, R},
	{0x2010, 0x2010, R},
	{0x2011, 0x2011, R},
	{0x2012, 0x2012, R},
	{0x2013, 0x2013, R},
	{0x2014, 0x2014, R},
	{0x2015, 0x2015, R},
	{0x2016, 0x2016, U},
	{0x2017, 0x2017, R},
	{0x2018, 0x2018, R},
	{0x2019, 0x2019, R},
	{0x201A, 0x201A, R},
	{0x201B, 0x201B, R},
	{0x201C, 0x201C, R},
	{0x201D, 0x201D, R},
	{0x201E, 0x201E, R},
	{0x201F, 0x201F, R},
	{0x2020, 0x2020, U},
	{0x2021, 0x2021, U},
	{0x2022, 0x2022, R},
	{0x2023, 0x2023, R},
	{0x2024, 0x2024, R},
	{0x2025, 0x2025, R},
	{0x2026, 0x2026, R},
	{0x2027, 0x2027, R},
	{0x2028, 0x2029, R},
	{0x202A, 0x202E, R},
	{0x202F, 0x202F, R},
	{0x2030, 0x2030, U},
	{0x2031, 0x2031, U},
	{0x2032, 0x2032, R},
	{0x2033, 0x2033, R},
	{0x2034, 0x2034, R},
	{0x2035, 0x2035, R},
	{0x2036, 0x2036, R},
	{0x2037, 0x2037, R},
	{0x2038, 0x2038, R},
	{0x2039, 0x2039, R},
	{0x203A, 0x203A, R},
	{0x203B, 0x203B, U},
	{0x203C, 0x203C, U},
	{0x203D, 0x203D, R},
	{0x203E, 0x203E, R},
	{0x203F, 0x203F, R},
	{0x2040, 0x2040, R},
	{0x2041, 0x2041, R},
	{0x2042, 0x2042, U},
	{0x2043, 0x2043, R},
	{0x2044, 0x2044, R},
	{0x2045, 0x2045, R},
	{0x2046, 0x2046, R},
	{0x2047, 0x2047, U},
	{0x2048, 0x2048, U},
	{0x2049, 0x2049, U},
	{0x204A, 0x204A, R},
	{0x204B, 0x204B, R},
	{0x204C, 0x204C, R},
	{0x204D, 0x204D, R},
	{0x204E, 0x204E, R},
	{0x204F, 0x204F, R},
	{0x2050, 0x2050, R},
	{0x2051, 0x2051, U},
	{0x2052, 0x2052, R},
	{0x2053, 0x2053, R},
	{0x2054, 0x2054, R},
	{0x2055, 0x2055, R},
	{0x2056, 0x2056, R},
	{0x2057, 0x2057, R},
	{0x2058, 0x2058, R},
	{0x2059, 0x2059, R},
	{0x205A, 0x205A, R},
	{0x205B, 0x205B, R},
	{0x205C, 0x205C, R},
	{0x205D, 0x205D, R},
	{0x205E, 0x205E, R},
	{0x205F, 0x205F, R},
	{0x2060, 0x2064, R},
	{0x2065, 0x2065, U},
	{0x2066, 0x2069, R},
	{0x206A, 0x206F, R},
	{0x2070, 0x209F, R},
	{0x20A0, 0x20AB, R},
	{0x20AC, 0x20AC, R},
	{0x20AD, 0x20CF, R},
	{0x20D0, 0x20DC, R},
	{0x20DD, 0x20E0, U},
	{0x20E1, 0x20E1, R},
	{0x20E2, 0x20E4, U},
	{0x20E5, 0x20FF, R},
	{0x2100, 0x2100, U},
	{0x2101, 0x2101, U},
	{0x2102, 0x2102, R},
	{0x2103, 0x2103, U},
	{0x2104, 0x2104, U},
	{0x2105, 0x2105, U},
	{0x2106, 0x2106, U},
	{0x2107, 0x2107, U},
	{0x2108, 0x2108, U},
	{0x2109, 0x2109, U},
	{0x210A, 0x210A, R},
	{0x210B, 0x210B, R},
	{0x210C, 0x210C, R},
	{0x210D, 0x210D, R},
	{0x210E, 0x210E, R},
	{0x210F, 0x210F, U},
	{0x2110, 0x2110, R},
	{0x2111, 0x2111, R},
	{0x2112, 0x2112, R},
	{0x2113, 0x2113, U},
	{0x2114, 0x2114, U},
	{0x2115, 0x2115, R},
	{0x2116, 0x2116, U},
	{0x2117, 0x2117, U},
	{0x2118, 0x2118, R},
	{0x2119, 0x2119, R},
	{0x211A, 0x211A, R},
	{0x211B, 0x211B, R},
	{0x211C, 0x211C, R},
	{0x211D, 0x211D, R},
	{0x211E, 0x211E, U},
	{0x211F, 0x211F, U},
	{0x2120, 0x2120, U},
	{0x2121, 0x2121, U},
	{0x2122, 0x2122, U},
	{0x2123, 0x2123, U},
	{0x2124, 0x2124, R},
	{0x2125, 0x2125, U},
	{0x2126, 0x2126, R},
	{0x2127, 0x2127, U},
	{0x2128, 0x2128, R},
	{0x2129, 0x2129, U},
	{0x212A, 0x212A, R},
	{0x212B, 0x212B, R},
	{0x212C, 0x212C, R},
	{0x212D, 0x212D, R},
	{0x212E, 0x212E, U},
	{0x212F, 0x212F, R},
	{0x2130, 0x2130, R},
	{0x2131, 0x2131, R},
	{0x2132, 0x2132, R},
	{0x2133, 0x2133, R},
	{0x2134, 0x2134, R},
	{0x2135, 0x2135, U},
	{0x2136, 0x2136, U},
	{0x2137, 0x2137, U},
	{0x2138, 0x2138, U},
	{0x2139, 0x2139, U},
	{0x213A, 0x213A, U},
	{0x213B, 0x213B, U},
	{0x213C, 0x213C, U},
	{0x213D, 0x213D, U},
	{0x213E, 0x213E, U},
	{0x213F, 0x213F, U},
	{0x2140, 0x2140, R},
	{0x2141, 0x2141, R},
	{0x2142, 0x2142, R},
	{0x2143, 0x2143, R},
	{0x2144, 0x2144, R},
	{0x2145, 0x2145, U},
	{0x2146, 0x2146, U},
	{0x2147, 0x2147, U},
	{0x2148, 0x2148, U},
	{0x2149, 0x2149, U},
	{0x214A, 0x214A, U},
	{0x214B, 0x214B, R},
	{0x214C, 0x214C, U},
	{0x214D, 0x214D, U},
	{0x214E, 0x214E, R},
	{0x214F, 0x214F, U},
	{0x2150, 0x2189, U},
	{0x218A, 0x218B, R},
	{0x218C, 0x218F, U},
	{0x2190, 0x21FF, R},
	{0x2200, 0x221D, R},
	{0x221E, 0x221E, U},
	{0x221F, 0x2233, R},
	{0x2234, 0x2235, U},
	{0x2236, 0x22FF, R},
	{0x2300, 0x2307, U},
	{0x2308, 0x230B, R},
	{0x230C, 0x231F, U},
	{0x2320, 0x2320, R},
	{0x2321, 0x2321, R},
	{0x2322, 0x2323, R},
	{0x2324, 0x2328, U},
	{0x2329, 0x2329, Tr},
	{0x232A, 0x232A, Tr},
	{0x232B, 0x232B, U},
	{0x232C, 0x237C, R},
	{0x237D, 0x239A, U},
	{0x239B, 0x23B3, R},
	{0x23B4, 0x23B6, R},
	{0x23B7, 0x23B9, R},
	{0x23BA, 0x23BD, R},
	{0x23BE, 0x23CD, U},
	{0x23CE, 0x23CE, R},
	{0x23CF, 0x23CF, U},
	{0x23D0, 0x23D0, R},
	{0x23D1, 0x23DB, U},
	{0x23DC, 0x23E1, R},
	{0x23E2, 0x23FF, U},
	{0x2400, 0x2422, U},
	{0x2423, 0x2423, R},
	{0x2424, 0x243F, U},
	{0x2440, 0x245F, U},
	{0x2460, 0x24FF, U},
	{0x2500, 0x257F, R},
	{0x2580, 0x259F, R},
	{0x25A0, 0x25FF, U},
	{0x2600, 0x2619, U},
	{0x261A, 0x261F, R},
	{0x2620, 0x26FF, U},
	{0x2700, 0x2767, U},
	{0x2768, 0x2775, R},
	{0x2776, 0x2793, U},
	{0x2794, 0x27BF, R},
	{0x27C0, 0x27C4, R},
	{0x27C5, 0x27C6, R},
	{0x27C7, 0x27E5, R},
	{0x27E6, 0x27EF, R},
	{0x27F0, 0x27FF, R},
	{0x2800, 0x28FF, R},
	{0x2900, 0x297F, R},
	{0x2980, 0x2982, R},
	{0x2983, 0x2998, R},
	{0x2999, 0x29D7, R},
	{0x29D8, 0x29DB, R},
	{0x29DC, 0x29FB, R},
	{0x29FC, 0x29FD, R},
	{0x29FE, 0x29FF, R},
	{0x2A00, 0x2AFF, R},
	{0x2B00, 0x2B11, R},
	{0x2B12, 0x2B2F, U},
	{0x2B30, 0x2B4F, R},
	{0x2B50, 0x2B59, U},
	{0x2B5A, 0x2BB7, R},
	{0x2BB8, 0x2BEB, U},
	{0x2BEC, 0x2BEF, R},
	{0x2BF0, 0x2BFF, U},
	{0x2C00, 0x2C5F, R},
	{0x2C60, 0x2C7F, R},
	{0x2C80, 0x2CFF, R},
	{0x2D00, 0x2D2F, R},
	{0x2D30, 0x2D7F, R},
	{0x2D80, 0x2DDF, R},
	{0x2DE0, 0x2DFF, R},
	{0x2E00, 0x2E16, R},
	{0x2E17, 0x2E17, R},
	{0x2E18, 0x2E19, R},
	{0x2E1A, 0x2E1A, R},
	{0x2E1B, 0x2E1F, R},
	{0x2E20, 0x2E21, R},
	{0x2E22, 0x2E25, R},
	{0x2E26, 0x2E29, R},
	{0x2E2A, 0x2E39, R},
	{0x2E3A, 0x2E3A, R},
	{0x2E3B, 0x2E3B, R},
	{0x2E3C, 0x2E44, R},
	{0x2E45, 0x2E7F, R},
	{0x2E80, 0x2EFF, U},
	{0x2F00, 0x2FDF, U},
	{0x2FE0, 0x2FEF, U},
	{0x2FF0, 0x2FFF, U},
	{0x3000, 0x3000, U},
	{0x3001, 0x3001, Tu},
	{0x3002, 0x3002, Tu},
	{0x3003, 0x3003, U},
	{0x3004, 0x3004, U},
	{0x3005, 0x3005, U},
	{0x3006, 0x3006, U},
	{0x3007, 0x3007, U},
	{0x3008, 0x3008, Tr},
	{0x3009, 0x3009, Tr},
	{0x300A, 0x300A, Tr},
	{0x300B, 0x300B, Tr},
	{0x300C, 0x300C, Tr},
	{0x300D, 0x300D, Tr},
	{0x300E, 0x300E, Tr},
	{0x300F, 0x300F, Tr},
	{0x3010, 0x3010, Tr},
	{0x3011, 0x3011, Tr},
	{0x3012, 0x3012, U},
	{0x3013, 0x3013, U},
	{0x3014, 0x3014, Tr},
	{0x3015, 0x3015, Tr},
	{0x3016, 0x3016, Tr},
	{0x3017, 0x3017, Tr},
	{0x3018, 0x3018, Tr},
	{0x3019, 0x3019, Tr},
	{0x301A, 0x301A, Tr},
	{0x301B, 0x301B, Tr},
	{0x301C, 0x301C, Tr},
	{0x301D, 0x301D, Tr},
	{0x301E, 0x301E, Tr},
	{0x301F, 0x301F, Tr},
	{0x3020, 0x3020, U},
	{0x3021, 0x3021, U},
	{0x3022, 0x3022, U},
	{0x3023, 0x3023, U},
	{0x3024, 0x3024, U},
	{0x3025, 0x3025, U},
	{0x3026, 0x3026, U},
	{0x3027, 0x3027, U},
	{0x3028, 0x3028, U},
	{0x3029, 0x3029, U},
	{0x302A, 0x302A, U},
	{0x302B, 0x302B, U},
	{0x302C, 0x302C, U},
	{0x302D, 0x302D, U},
	{0x302E, 0x302E, U},
	{0x302F, 0x302F, U},
	{0x3030, 0x3030, Tr},
	{0x3031, 0x3031, U},
	{0x3032, 0x3032, U},
	{0x3033, 0x3033, U},
	{0x3034, 0x3034, U},
	{0x3035, 0x3035, U},
	{0x3036, 0x3036, U},
	{0x3037, 0x3037, U},
	{0x3038, 0x3038, U},
	{0x3039, 0x3039, U},
	{0x303A, 0x303A, U},
	{0x303B, 0x303B, U},
	{0x303C, 0x303C, U},
	{0x303D, 0x303D, U},
	{0x303E, 0x303E, U},
	{0x303F, 0x303F, U},
	{0x3040, 0x3040, U},
	{0x3041, 0x3041, Tu},
	{0x3042, 0x3042, U},
	{0x3043, 0x3043, Tu},
	{0x3044, 0x3044, U},
	{0x3045, 0x3045, Tu},
	{0x3046, 0x3046, U},
	{0x3047, 0x3047, Tu},
	{0x3048, 0x3048, U},
	{0x3049, 0x3049, Tu},
	{0x304A, 0x304A, U},
	{0x304B, 0x304B, U},
	{0x304C, 0x304C, U},
	{0x304D, 0x304D, U},
	{0x304E, 0x304E, U},
	{0x304F, 0x304F, U},
	{0x3050, 0x3050, U},
	{0x3051, 0x3051, U},
	{0x3052, 0x3052, U},
	{0x3053, 0x3053, U},
	{0x3054, 0x3054, U},
	{0x3055, 0x3055, U},
	{0x3056, 0x3056, U},
	{0x3057, 0x3057, U},
	{0x3058, 0x3058, U},
	{0x3059, 0x3059, U},
	{0x305A, 0x305A, U},
	{0x305B, 0x305B, U},
	{0x305C, 0x305C, U},
	{0x305D, 0x305D, U},
	{0x305E, 0x305E, U},
	{0x305F, 0x305F, U},
	{0x3060, 0x3060, U},
	{0x3061, 0x3061, U},
	{0x3062, 0x3062, U},
	{0x3063, 0x3063, Tu},
	{0x3064, 0x3064, U},
	{0x3065, 0x3065, U},
	{0x3066, 0x3066, U},
	{0x3067, 0x3067, U},
	{0x3068, 0x3068, U},
	{0x3069, 0x3069, U},
	{0x306A, 0x306A, U},
	{0x306B, 0x306B, U},
	{0x306C, 0x306C, U},
	{0x306D, 0x306D, U},
	{0x306E, 0x306E, U},
	{0x306F, 0x306F, U},
	{0x3070, 0x3070, U},
	{0x3071, 0x3071, U},
	{0x3072, 0x3072, U},
	{0x3073, 0x3073, U},
	{0x3074, 0x3074, U},
	{0x3075, 0x3075, U},
	{0x3076, 0x3076, U},
	{0x3077, 0x3077, U},
	{0x3078, 0x3078, U},
	{0x3079, 0x3079, U},
	{0x307A, 0x307A, U},
	{0x307B, 0x307B, U},
	{0x307C, 0x307C, U},
	{0x307D, 0x307D, U},
	{0x307E, 0x307E, U},
	{0x307F, 0x307F, U},
	{0x3080, 0x3080, U},
	{0x3081, 0x3081, U},
	{0x3082, 0x3082, U},
	{0x3083, 0x3083, Tu},
	{0x3084, 0x3084, U},
	{0x3085, 0x3085, Tu},
	{0x3086, 0x3086, U},
	{0x3087, 0x3087, Tu},
	{0x3088, 0x3088, U},
	{0x3089, 0x3089, U},
	{0x308A, 0x308A, U},
	{0x308B, 0x308B, U},
	{0x308C, 0x308C, U},
	{0x308D, 0x308D, U},
	{0x308E, 0x308E, Tu},
	{0x308F, 0x308F, U},
	{0x3090, 0x3090, U},
	{0x3091, 0x3091, U},
	{0x3092, 0x3092, U},
	{0x3093, 0x3093, U},
	{0x3094, 0x3094, U},
	{0x3095, 0x3095, Tu},
	{0x3096, 0x3096, Tu},
	{0x3097, 0x3097, U},
	{0x3098, 0x3098, U},
	{0x3099, 0x3099, U},
	{0x309A, 0x309A, U},
	{0x309B, 0x309B, Tu},
	{0x309C, 0x309C, Tu},
	{0x309D, 0x309D, U},
	{0x309E, 0x309E, U},
	{0x309F, 0x309F, U},
	{0x30A0, 0x30A0, Tr},
	{0x30A1, 0x30A1, Tu},
	{0x30A2, 0x30A2, U},
	{0x30A3, 0x30A3, Tu},
	{0x30A4, 0x30A4, U},
	{0x30A5, 0x30A5, Tu},
	{0x30A6, 0x30A6, U},
	{0x30A7, 0x30A7, Tu},
	{0x30A8, 0x30A8, U},
	{0x30A9, 0x30A9, Tu},
	{0x30AA, 0x30AA, U},
	{0x30AB, 0x30AB, U},
	{0x30AC, 0x30AC, U},
	{0x30AD, 0x30AD, U},
	{0x30AE, 0x30AE, U},
	{0x30AF, 0x30AF, U},
	{0x30B0, 0x30B0, U},
	{0x30B1, 0x30B1, U},
	{0x30B2, 0x30B2, U},
	{0x30B3, 0x30B3, U},
	{0x30B4, 0x30B4, U},
	{0x30B5, 0x30B5, U},
	{0x30B6, 0x30B6, U},
	{0x30B7, 0x30B7, U},
	{0x30B8, 0x30B8, U},
	{0x30B9, 0x30B9, U},
	{0x30BA, 0x30BA, U},
	{0x30BB, 0x30BB, U},
	{0x30BC, 0x30BC, U},
	{0x30BD, 0x30BD, U},
	{0x30BE, 0x30BE, U},
	{0x30BF, 0x30BF, U},
	{0x30C0, 0x30C0, U},
	{0x30C1, 0x30C1, U},
	{0x30C2, 0x30C2, U},
	{0x30C3, 0x30C3, Tu},
	{0x30C4, 0x30C4, U},
	{0x30C5, 0x30C5, U},
	{0x30C6, 0x30C6, U},
	{0x30C7, 0x30C7, U},
	{0x30C8, 0x30C8, U},
	{0x30C9, 0x30C9, U},
	{0x30CA, 0x30CA, U},
	{0x30CB, 0x30CB, U},
	{0x30CC, 0x30CC, U},
	{0x30CD, 0x30CD, U},
	{0x30CE, 0x30CE, U},
	{0x30CF, 0x30CF, U},
	{0x30D0, 0x30D0, U},
	{0x30D1, 0x30D1, U},
	{0x30D2, 0x30D2, U},
	{0x30D3, 0x30D3, U},
	{0x30D4, 0x30D4, U},
	{0x30D5, 0x30D5, U},
	{0x30D6, 0x30D6, U},
	{0x30D7, 0x30D7, U},
	{0x30D8, 0x30D8, U},
	{0x30D9, 0x30D9, U},
	{0x30DA, 0x30DA, U},
	{0x30DB, 0x30DB, U},
	{0x30DC, 0x30DC, U},
	{0x30DD, 0x30DD, U},
	{0x30DE, 0x30DE, U},
	{0x30DF, 0x30DF, U},
	{0x30E0, 0x30E0, U},
	{0x30E1, 0x30E1, U},
	{0x30E2, 0x30E2, U},
	{0x30E3, 0x30E3, Tu},
	{0x30E4, 0x30E4, U},
	{0x30E5, 0x30E5, Tu},
	{0x30E6, 0x30E6, U},
	{0x30E7, 0x30E7, Tu},
	{0x30E8, 0x30E8, U},
	{0x30E9, 0x30E9, U},
	{0x30EA, 0x30EA, U},
	{0x30EB, 0x30EB, U},
	{0x30EC, 0x30EC, U},
	{0x30ED, 0x30ED, U},
	{0x30EE, 0x30EE, Tu},
	{0x30EF, 0x30EF, U},
	{0x30F0, 0x30F0, U},
	{0x30F1, 0x30F1, U},
	{0x30F2, 0x30F2, U},
	{0x30F3, 0x30F3, U},
	{0x30F4, 0x30F4, U},
	{0x30F5, 0x30F5, Tu},
	{0x30F6, 0x30F6, Tu},
	{0x30F7, 0x30F7, U},
	{0x30F8, 0x30F8, U},
	{0x30F9, 0x30F9, U},
	{0x30FA, 0x30FA, U},
	{0x30FB, 0x30FB, U},
	{0x30FC, 0x30FC, Tr},
	{0x30FD, 0x30FD, U},
	{0x30FE, 0x30FE, U},
	{0x30FF, 0x30FF, U},
	{0x3100, 0x3126, U},
	{0x3127, 0x3127, Tu},
	{0x3128, 0x312F, U},
	{0x3130, 0x318F, U},
	{0x3190, 0x319F, U},
	{0x31A0, 0x31BF, U},
	{0x31C0, 0x31EF, U},
	{0x31F0, 0x31FF, Tu},
	{0x3200, 0x321E, U},
	{0x321F, 0x32FF, U},
	{0x3300, 0x3357, Tu},
	{0x3358, 0x337A, U},
	{0x337B, 0x337F, Tu},
	{0x3380, 0x33FF, U},
	{0x3400, 0x4DBF, U},
	{0x4DC0, 0x4DFF, U},
	{0x4E00, 0x9FFF, U},
	{0xA000, 0xA48F, U},
	{0xA490, 0xA4CF, U},
	{0xA4D0, 0xA4FF, R},
	{0xA500, 0xA63F, R},
	{0xA640, 0xA69F, R},
	{0xA6A0, 0xA6FF, R},
	{0xA700, 0xA71F, R},
	{0xA720, 0xA7FF, R},
	{0xA800, 0xA82F, R},
	{0xA830, 0xA83F, R},
	{0xA840, 0xA87F, R},
	{0xA880, 0xA8DF, R},
	{0xA8E0, 0xA8FF, R},
	{0xA900, 0xA92F, R},
	{0xA930, 0xA95F, R},
	{0xA960, 0xA97F, U},
	{0xA980, 0xA9DF, R},
	{0xA9E0, 0xA9FF, R},
	{0xAA00, 0xAA5F, R},
	{0xAA60, 0xAA7F, R},
	{0xAA80, 0xAADF, R},
	{0xAAE0, 0xAAFF, R},
	{0xAB00, 0xAB2F, R},
	{0xAB30, 0xAB6F, R},
	{0xAB70, 0xABBF, R},
	{0xABC0, 0xABFF, R},
	{0xAC00, 0xD7AF, U},
	{0xD7B0, 0xD7FF, U},
	{0xD800, 0xDFFF, R},
	{0xE000, 0xF8FF, U},
	{0xF900, 0xFAFF, U},
	{0xFB00, 0xFB4F, R},
	{0xFB50, 0xFDFF, R},
	{0xFE00, 0xFE0F, R},
	{0xFE10, 0xFE10, U},
	{0xFE11, 0xFE11, U},
	{0xFE12, 0xFE12, U},
	{0xFE13, 0xFE13, U},
	{0xFE14, 0xFE14, U},
	{0xFE15, 0xFE15, U},
	{0xFE16, 0xFE16, U},
	{0xFE17, 0xFE17, U},
	{0xFE18, 0xFE18, U},
	{0xFE19, 0xFE19, U},
	{0xFE1A, 0xFE1F, U},
	{0xFE20, 0xFE2F, R},
	{0xFE30, 0xFE30, U},
	{0xFE31, 0xFE31, U},
	{0xFE32, 0xFE32, U},
	{0xFE33, 0xFE33, U},
	{0xFE34, 0xFE34, U},
	{0xFE35, 0xFE35, U},
	{0xFE36, 0xFE36, U},
	{0xFE37, 0xFE37, U},
	{0xFE38, 0xFE38, U},
	{0xFE39, 0xFE39, U},
	{0xFE3A, 0xFE3A, U},
	{0xFE3B, 0xFE3B, U},
	{0xFE3C, 0xFE3C, U},
	{0xFE3D, 0xFE3D, U},
	{0xFE3E, 0xFE3E, U},
	{0xFE3F, 0xFE3F, U},
	{0xFE40, 0xFE40, U},
	{0xFE41, 0xFE41, U},
	{0xFE42, 0xFE42, U},
	{0xFE43, 0xFE43, U},
	{0xFE44, 0xFE44, U},
	{0xFE45, 0xFE45, U},
	{0xFE46, 0xFE46, U},
	{0xFE47, 0xFE47, U},
	{0xFE48, 0xFE48, U},
	{0xFE49, 0xFE49, R},
	{0xFE4A, 0xFE4A, R},
	{0xFE4B, 0xFE4B, R},
	{0xFE4C, 0xFE4C, R},
	{0xFE4D, 0xFE4D, R},
	{0xFE4E, 0xFE4E, R},
	{0xFE4F, 0xFE4F, R},
	{0xFE50, 0xFE50, Tu},
	{0xFE51, 0xFE51, Tu},
	{0xFE52, 0xFE52, Tu},
	{0xFE53, 0xFE53, U},
	{0xFE54, 0xFE54, U},
	{0xFE55, 0xFE55, U},
	{0xFE56, 0xFE56, U},
	{0xFE57, 0xFE57, U},
	{0xFE58, 0xFE58, R},
	{0xFE59, 0xFE59, Tr},
	{0xFE5A, 0xFE5A, Tr},
	{0xFE5B, 0xFE5B, Tr},
	{0xFE5C, 0xFE5C, Tr},
	{0xFE5D, 0xFE5D, Tr},
	{0xFE5E, 0xFE5E, Tr},
	{0xFE5F, 0xFE5F, U},
	{0xFE60, 0xFE60, U},
	{0xFE61, 0xFE61, U},
	{0xFE62, 0xFE62, U},
	{0xFE63, 0xFE63, R},
	{0xFE64, 0xFE64, R},
	{0xFE65, 0xFE65, R},
	{0xFE66, 0xFE66, R},
	{0xFE67, 0xFE67, U},
	{0xFE68, 0xFE68, U},
	{0xFE69, 0xFE69, U},
	{0xFE6A, 0xFE6A, U},
	{0xFE6B, 0xFE6B, U},
	{0xFE6C, 0xFE6F, U},
	{0xFE70, 0xFEFF, R},
	{0xFF00, 0xFF00, R},
	{0xFF01, 0xFF01, Tu},
	{0xFF02, 0xFF02, U},
	{0xFF03, 0xFF03, U},
	{0xFF04, 0xFF04, U},
	{0xFF05, 0xFF05, U},
	{0xFF06, 0xFF06, U},
	{0xFF07, 0xFF07, U},
	{0xFF08, 0xFF08, Tr},
	{0xFF09, 0xFF09, Tr},
	{0xFF0A, 0xFF0A, U},
	{0xFF0B, 0xFF0B, U},
	{0xFF0C, 0xFF0C, Tu},
	{0xFF0D, 0xFF0D, R},
	{0xFF0E, 0xFF0E, Tu},
	{0xFF0F, 0xFF0F, U},
	{0xFF10, 0xFF10, U},
	{0xFF11, 0xFF11, U},
	{0xFF12, 0xFF12, U},
	{0xFF13, 0xFF13, U},
	{0xFF14, 0xFF14, U},
	{0xFF15, 0xFF15, U},
	{0xFF16, 0xFF16, U},
	{0xFF17, 0xFF17, U},
	{0xFF18, 0xFF18, U},
	{0xFF19, 0xFF19, U},
	{0xFF1A, 0xFF1A, Tr},
	{0xFF1B, 0xFF1B, Tr},
	{0xFF1C, 0xFF1C, R},
	{0xFF1D, 0xFF1D, R},
	{0xFF1E, 0xFF1E, R},
	{0xFF1F, 0xFF1F, Tu},
	{0xFF20, 0xFF20, U},
	{0xFF21, 0xFF3A, U},
	{0xFF3B, 0xFF3B, Tr},
	{0xFF3C, 0xFF3C, U},
	{0xFF3D, 0xFF3D, Tr},
	{0xFF3E, 0xFF3E, U},
	{0xFF3F, 0xFF3F, Tr},
	{0xFF40, 0xFF40, U},
	{0xFF41, 0xFF5A, U},
	{0xFF5B, 0xFF5B, Tr},
	{0xFF5C, 0xFF5C, Tr},
	{0xFF5D, 0xFF5D, Tr},
	{0xFF5E, 0xFF5E, Tr},
	{0xFF5F, 0xFF5F, Tr},
	{0xFF60, 0xFF60, Tr},
	{0xFF61, 0xFF61, R},
	{0xFF62, 0xFF62, R},
	{0xFF63, 0xFF63, R},
	{0xFF64, 0xFF64, R},
	{0xFF65, 0xFF65, R},
	{0xFF66, 0xFF66, R},
	{0xFF67, 0xFF67, R},
	{0xFF68, 0xFF68, R},
	{0xFF69, 0xFF69, R},
	{0xFF6A, 0xFF6A, R},
	{0xFF6B, 0xFF6B, R},
	{0xFF6C, 0xFF6C, R},
	{0xFF6D, 0xFF6D, R},
	{0xFF6E, 0xFF6E, R},
	{0xFF6F, 0xFF6F, R},
	{0xFF70, 0xFF70, R},
	{0xFF71, 0xFF9F, R},
	{0xFFA0, 0xFFDF, R},
	{0xFFE0, 0xFFE0, U},
	{0xFFE1, 0xFFE1, U},
	{0xFFE2, 0xFFE2, U},
	{0xFFE3, 0xFFE3, Tr},
	{0xFFE4, 0xFFE4, U},
	{0xFFE5, 0xFFE5, U},
	{0xFFE6, 0xFFE6, U},
	{0xFFE7, 0xFFE7, U},
	{0xFFE8, 0xFFE8, R},
	{0xFFE9, 0xFFE9, R},
	{0xFFEA, 0xFFEA, R},
	{0xFFEB, 0xFFEB, R},
	{0xFFEC, 0xFFEC, R},
	{0xFFED, 0xFFED, R},
	{0xFFEE, 0xFFEE, R},
	{0xFFEF, 0xFFEF, R},
	{0xFFF0, 0xFFF8, U},
	{0xFFF9, 0xFFFB, R},
	{0xFFFC, 0xFFFC, U},
	{0xFFFD, 0xFFFD, U},
	{0xFFFE, 0xFFFE, R},
	{0xFFFF, 0xFFFF, R},
	{0x10000, 0x1007F, R},
	{0x10080, 0x100FF, R},
	{0x10100, 0x1013F, R},
	{0x10140, 0x1018F, R},
	{0x10190, 0x101CF, R},
	{0x101D0, 0x101FF, R},
	{0x10280, 0x1029F, R},
	{0x102A0, 0x102DF, R},
	{0x102E0, 0x102FF, R},
	{0x10300, 0x1032F, R},
	{0x10330, 0x1034F, R},
	{0x10350, 0x1037F, R},
	{0x10380, 0x1039F, R},
	{0x103A0, 0x103DF, R},
	{0x10400, 0x1044F, R},
	{0x10450, 0x1047F, R},
	{0x10480, 0x104AF, R},
	{0x104B0, 0x104FF, R},
	{0x10500, 0x1052F, R},
	{0x10530, 0x1056F, R},
	{0x10600, 0x1077F, R},
	{0x10800, 0x1083F, R},
	{0x10840, 0x1085F, R},
	{0x10860, 0x1087F, R},
	{0x10880, 0x108AF, R},
	{0x108E0, 0x108FF, R},
	{0x10900, 0x1091F, R},
	{0x10920, 0x1093F, R},
	{0x10980, 0x1099F, U},
	{0x109A0, 0x109FF, R},
	{0x10A00, 0x10A5F, R},
	{0x10A60, 0x10A7F, R},
	{0x10A80, 0x10A9F, R},
	{0x10AC0, 0x10AFF, R},
	{0x10B00, 0x10B3F, R},
	{0x10B40, 0x10B5F, R},
	{0x10B60, 0x10B7F, R},
	{0x10B80, 0x10BAF, R},
	{0x10C00, 0x10C4F, R},
	{0x10C80, 0x10CFF, R},
	{0x10E60, 0x10E7F, R},
	{0x11000, 0x1107F, R},
	{0x11080, 0x110CF, R},
	{0x110D0, 0x110FF, R},
	{0x11100, 0x1114F, R},
	{0x11150, 0x1117F, R},
	{0x11180, 0x111DF, R},
	{0x111E0, 0x111FF, R},
	{0x11200, 0x1124F, R},
	{0x11280, 0x112AF, R},
	{0x112B0, 0x112FF, R},
	{0x11300, 0x1137F, R},
	{0x11400, 0x1147F, R},
	{0x11480, 0x114DF, R},
	{0x11580, 0x115FF, U},
	{0x11600, 0x1165F, R},
	{0x11660, 0x1167F, R},
	{0x11680, 0x116CF, R},
	{0x11700, 0x1173F, R},
	{0x118A0, 0x118FF, R},
	{0x11AC0, 0x11AFF, R},
	{0x11C00, 0x11C6F, R},
	{0x11C70, 0x11CBF, R},
	{0x12000, 0x123FF, R},
	{0x12400, 0x1247F, R},
	{0x12480, 0x1254F, R},
	{0x13000, 0x1342F, U},
	{0x14400, 0x1467F, U},
	{0x16800, 0x16A3F, R},
	{0x16A40, 0x16A6F, R},
	{0x16AD0, 0x16AFF, R},
	{0x16B00, 0x16B8F, R},
	{0x16F00, 0x16F9F, R},
	{0x16FE0, 0x16FFF, U},
	{0x17000, 0x187FF, U},
	{0x18800, 0x18AFF, U},
	{0x1B000, 0x1B0FF, U},
	{0x1BC00, 0x1BC9F, R},
	{0x1BCA0, 0x1BCAF, R},
	{0x1D000, 0x1D0FF, U},
	{0x1D100, 0x1D1FF, U},
	{0x1D200, 0x1D24F, R},
	{0x1D300, 0x1D35F, U},
	{0x1D360, 0x1D37F, U},
	{0x1D400, 0x1D7FF, R},
	{0x1D800, 0x1DAAF, U},
	{0x1E000, 0x1E02F, R},
	{0x1E800, 0x1E8DF, R},
	{0x1E900, 0x1E95F, R},
	{0x1EE00, 0x1EEFF, R},
	{0x1F000, 0x1F02F, U},
	{0x1F030, 0x1F09F, U},
	{0x1F0A0, 0x1F0FF, U},
	{0x1F100, 0x1F1FF, U},
	{0x1F200, 0x1F200, Tu},
	{0x1F201, 0x1F201, Tu},
	{0x1F202, 0x1F2FF, U},
	{0x1F300, 0x1F5FF, U},
	{0x1F600, 0x1F64F, U},
	{0x1F650, 0x1F67F, U},
	{0x1F680, 0x1F6FF, U},
	{0x1F700, 0x1F77F, U},
	{0x1F780, 0x1F7FF, U},
	{0x1F800, 0x1F8FF, R},
	{0x1F900, 0x1F9FF, U},
	{0x20000, 0x2A6DF, U},
	{0x2A6E0, 0x2A6FF, U},
	{0x2A700, 0x2B73F, U},
	{0x2B740, 0x2B81F, U},
	{0x2B820, 0x2CEAF, U},
	{0x2CEB0, 0x2F7FF, U},
	{0x2F800, 0x2FA1F, U},
	{0x2FA20, 0x2FFFD, U},
	{0x30000, 0x3FFFD, U},
	{0xE0000, 0xE007F, R},
	{0xE0100, 0xE01EF, R},
	{0xF0000, 0xFFFFD, U},
	{0x100000, 0x10FFFD, U},
}
