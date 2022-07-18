
FMT CHEAT SHEET

GENERAL
====

1. %v (value in default format)
2. %T(Type)
3. %% (literal %)

BOOLEAN
==
1. %t (true or false)

INTEGER
==
1. %b (base2)
2. %o (base8)
3. %d (base10)
4. %x (base16)

FLAOTING POINTS
==
1. %e(scientific notation)
2. %f/%F (decimal or exponent)
3. %g (for large exponent)

STRINGS
==
1. %s (default)
2. %q (double quoted strings)

WIDTH AND PERCISION
==
1. %f (default Width, default precision)
2. %9f ( Width 9, default precision)
3. %.2f (default Width, precision 2)
4. %9.29f (Width 9, precision 2)
5. %9.f (Width 9, precision 0)

PADDING
==
1. %09d (pads digit to length 9 with preceeding 0's)
2. %-4d (pads with spaces (Width 4 left justified))

METHODS
==
1. Sprintf() format without printing
2. Printf() format with printing