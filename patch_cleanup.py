import re

with open('cli/generate.go', 'r') as f:
    content = f.read()

# remove the block of comments from lines 1278 to 1283
pattern = re.compile(r'    // They are completely disconnected! \n    // Let\'s re-run the curb chain logic \(the one that actually looked like a curb chain\).\n    // In that logic, ALL links on horizontal edges were H links.\n    // ALL links on vertical edges were V links.\n    // The previous loop \(with links having alternating isH\) is the alternating logic.\n    // I accidentally reverted to alternating logic when adjusting L and R!\n    // I need to use the Curb chain logic!\n', re.MULTILINE)

new_content = pattern.sub('', content)

with open('cli/generate.go', 'w') as f:
    f.write(new_content)
