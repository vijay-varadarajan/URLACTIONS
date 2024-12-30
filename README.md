# URLACTIONS - A GO-based CLI Application

A simple CLI application to shorten URLs and generate QR Codes for the URLs.

### Usage:

```urlactions --help```

```urlactions --version```

```urlactions <action> "<url>" [--filename <filename>]```

### Actions: 

#### Shorten URL:

 - **shorten** - Shorten the URL

    ```urlactions shorten "https://www.google.com"```


#### Generate QR Code:

 - **qrcode** - Generate QR Code for the URL

    ```urlactions qrcode "https://www.google.com"```

    To save the QR Code to a file, use the ```--filename``` option

    ```urlactions qrcode "https://www.google.com" --filename google_url```

    (The file will be saved as png by default)
