# IDN-CHEKING

Check available places to make an appointment at portaal.refugeepass.nl
The program sends a request every 30 seconds and checks the availability of free slots in the portal across the country (NL)


## Usage

Go tothe link https://portaal.refugeepass.nl/en/make-an-appointment, press F12 and try to find an appointment. 
If there are no slots, take cookie and a token from the site (Debug console -> Network -> XHD -> select last query -> Header).
Then put the data in the `config.toml` file.




```
$ make build
$ make run

time="11-08-2022 12:53:41" level=info msg="make request on 2022-08-17 ..."
time="11-08-2022 12:53:48" level=info msg="Responce: []"
time="11-08-2022 12:53:48" level=info msg="Nothing found"
time="11-08-2022 12:53:48" level=info msg=sleep...

```