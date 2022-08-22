func reverseBits(num uint32) uint32 {
    var output uint32
    for i := 0; i < 32; i++ {
        output, num = (output << 1) + (num % 2), num >> 1
    }
    return output
}
