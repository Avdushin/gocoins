![gocoins](logo.png)

# GOCOINS

#### Widget POLYBAR for monitoring the exchange rate against the ruble.

### USAGE:

POLYBAR script `~/.config/polybar/config`

```text
[module/course]
type = custom/script
interval = 60
exec = yuan; euro
label-foreground = ${colors.foreground}
format-background = 
format-underline = #7D49B6
```

Output your module:

```text
modules-left =   
modules-center = 
modules-right =  course
```



You can edit `exec = currency` to change output!



P.S this program copy currencys to the `/usr/bin` to do their bynary...



