# Simple Go HTTP App Generator

Generates a simple http server using go HTTP package.


## Install the template

```sh
$ ironman install https://github.com/ironman-project/simple-gohttp-template.git
```


## Generators

### App

 ```sh
 $ ironman generate simple-gohttp /path/to/app --set projectName="Some project name",projectDescription="Some project description"
 ```

 #### Values

 * projectName: A project name
 * projectDescription: A description of the project
 * myName: Name to display in the root enpoint with the "Hello {{.myName}}" message.


