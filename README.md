# URLACTIONS

A simple command line tool to perform various actions on URLs. 


### Setup:
1. Download the latest release from the [Releases](https://github.com/vijay-varadarajan/urlactions/releases)
2. Extract the contents of the zip file

##### For macOS and Linux:
3. Open the terminal and navigate to the extracted folder
4. Run the following command to install the tool:

   ```bash
   chmod +x install.sh
   ./install.sh
   ```

##### For Windows:
3. Open Command Prompt and navigate to the extracted folder
4. Run the following command to install the tool:

   ```cmd
   install.bat
   ```

5. The tool is now installed and ready to use


### Usage:

```urlactions -<action> "<url>" [-save [-filename <filename>]]```

*If macOS: Use __urlcli__ instead of __urlactions__ command*


### Actions: 


#### Validate URL

 - ```-validate``` - Validate the URL

    Example: ```urlactions -validate "https://www.google.com"```

    Checks if URL is valid and if server is reachable.
    The validation result is displayed on the console.


#### Shorten URL:

 - ```-shorten``` - Shorten the URL

    Example: ```urlactions -shorten "https://www.google.com"```

    Shortened URLs cannot be saved to a file.
    They are copied to the clipboard automatically.


#### Generate QR Code:

 - ```-qrcode``` - Generate QR Code for the URL

    Example: ```urlactions qrcode "https://www.google.com"```

    To save the QR Code to a file with default name, use the ```-save``` option

    Example: ```urlactions qrcode "https://www.google.com" -save```

    (The file will be saved as domain_name.png by default)

    To save the QR Code to a file with a custom name, use the ```-save``` and ```-filename``` options

    Example: ```urlactions qrcode "https://www.google.com" -save -filename "google"```

    (The file will be saved as google.png)

    Files are saved in **PNG format** by default.


#### Version:

 - ```-version``` - Display the version of the application

    Example: ```urlactions -version```


#### Help:

 - ```-help``` - Display the help message

    Example: ```urlactions -help```