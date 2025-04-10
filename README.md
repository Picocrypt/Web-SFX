# Picocrypt SFX
Convert a standard .pcv into a self-contained, self-decrypting/extracting .html file for ultimate portability!

# Installation
If you don't have Go installed, download it from <a href="https://go.dev/dl/">here</a> or install it from your package manager. Next:
```
go install github.com/Picocrypt/Web-SFX/picocrypt-sfx@latest
```
You should now be able to run `picocrypt-sfx`. If not, run `export PATH=$PATH:$(go env GOPATH)/bin` and try again.

If you don't want to install Go, compiled executables are available in Releases.

# Usage
Only use this tool on volumes smaller than 1 GiB that use no advanced features or keyfiles.
```
picocrypt-sfx secret.txt.pcv
```
A `secret.txt.pcv.html` will be created which embeds the volume. **Always test the .html file after creation!**

# Use Cases
- **Secure file sharing**: easily share encrypted files to others without needing them to download software (that they may not trust or know how to use) on their end.
- **Long-term file storage**: the decryption module is bundled with the volume data into a single .html file, meaning everything you need to access your encrypted files is present and self-contained (works offline). Due to the stable and standardized nature of the web and WebAssembly, this single .html file will be able to decrypt your files on any platform with any modern browser indefinitely unless web browsers stop existing.

# Issues
Report any issues to the main Picocrypt/Picocrypt repository and prefix the title with 'Web-SFX:'.
