# Password Generator
CLI to generate cryptographically secure passwords in your local terminal.

## Install

    go get https://github.com/chiips/password-generator

## Use

The `generate` command (aliases: `gen`, `g`) asks your requirements and generates a password accordingly.

Choose whether you need uppercase letters, lowercase letters, numbers, or special characters, and how long your password needs to be.

Set `-language=francais` or `-language=italiano` (aliases: `lang`, `l`) for French or Italian versions. The default language is English.

Examples:  
`password generate`  
`password gen -l francais`  

## Licence

This code is licensed under the MIT License.