# Picocrypt SFX
Convert a standard .pcv into a self-contained, self-decrypting/extracting .html file for ultimate portability!

# Installation
If you don't have Go installed, download it from <a href="https://go.dev/dl/">here</a> or install it from your package manager. Next:
```
go install github.com/Picocrypt/Web-SFX/picocrypt-sfx@latest
```
You should now be able to run `picocrypt-sfx`. If not, run `export PATH=$PATH:$(go env GOPATH)/bin` and try again.

# Usage
Only use this tool on volumes smaller than 1 GiB that use no advanced features nor keyfiles.
```
picocrypt-sfx secret.txt.pcv
```
A `secret.txt.pcv.html` will be created which embeds the volume. **Always test the .html file after creation!**

# Use Cases
- **Long-term file storage**: because the decryption code is now bundled with the volume data into a single .html file, everything you need to access your encrypted files is present and self-contained. Due to the web and WebAssembly being standardized and stable, this single .html file will be able to decrypt your files on any platform with any modern browser indefinitely far into the future. This is arguably one of the safest and most reliable long-term encryption solutions as there's no external dependencies at all.
- **File sharing**: easily share encrypted files to others without having them download software (that they may not trust or know how to use).
