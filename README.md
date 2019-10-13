# yam
Yara rules fetcher and merger - get rules from (exclude mibile rules) https://github.com/Yara-Rules/rules.

Results **rules.yara** file can be used with **volatility  yarascan** plugin.

#### Usage
To use just run
```cmd
yam.exe
```
It will create **rules.yara ** file in current folder.

Next use it with volatility like in this example
```cmd
volatility -f dump.raw --profile=WinXPSP2x86 yarascan -y rules.yar
```
