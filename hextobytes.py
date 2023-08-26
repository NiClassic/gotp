import sys

def hextobytes(hex):
    chunks = [hex[i:i+2] for i in range(0, len(hex), 2)]
    bytes = [int(x, 16) for x in chunks]
    return bytes
    
print(f'{hextobytes(sys.argv[1])}')