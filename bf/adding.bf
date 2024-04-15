[
    simple adding test/example
]

+++++++++++++++++++++++++++++++++ set the first byte to 33("!" in ASCII)

. print second byte

> move to second byte
++++++++++++++++++++++++++++++++++++++++ set the second byte to 40("(" in ASCII)

. print second byte

< move back to first byte


add first byte to second byte giving 73 which is an uppercase "i" in ASCII
[ is the current byte zero? jump to corresponding bracket if so
    - remove 1 from current byte
    > move to next byte
    + add 1 to current byte
    < move to previous byte
] is the current byte non zero? jump to corresponding bracket if so

> move to second byte

. print second byte