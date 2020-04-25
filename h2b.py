import binascii
import sys

inp = sys.argv[1]
out = sys.argv[2]
with open(inp) as f, open(out, 'wb') as fout:
    for line in f:
        fout.write(
            binascii.unhexlify(''.join(line.split()))
        )