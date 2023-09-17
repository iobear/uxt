# uxt

Unix time helper command

uxt<br />
→946684800

uxt +{int} or -{int}<br />
→946684810

uxt {Unix Epoch}<br />
→Thu May  8 08:13:30 CEST 2025

uxt {Unix Epoch} RFC3339<br />
→2025-05-08T08:13:30+02:00

uxt since {Unix Epoch}<br />
→60 days, 18 hours, 5 minutes, 44 seconds

uxt serv:< port > (local webserver)
```
/current
/plus/<*int>
/minus/<*int>
/since/<*uxt>
/rfc3389/<*uxt>
```

uxt version<br />
→0.0.x

![alt uxt](docs/uxt.gif?raw=true "uxt demo")