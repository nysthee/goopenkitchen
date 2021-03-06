# goopenkitchen

**Summary**

This training on the go language describes the language from ground up using examples.
Most examples can be run and edited from within the presentation.

---

**Viewing a hosted version of the presentation in your browser** (=preferred approach)

The interactive presentation is available via the following link:

http://go-talks.appspot.com/github.com/MarcGrol/goopenkitchen/openKitchen.slide#1

---

**Viewing it locally using the "present"-tool**

You can also run it locally on your own device:
Use the "present"-tool to "run" the presentation.

*Step 1 - Get the present tool*

    go get golang.org/x/tools/present
    # present binary should now be in ${GOPATH}/bin 

*Step 2 - Fetch the presentation from github*

    go get github.com/MarcGrol/goopenkitchen
    or
    mkdir -p ${GOPATH}/src/github.com/MarcGrol
    cd ${GOPATH}/src/github.com/MarcGrol
    git clone https://github.com/MarcGrol/goopenkitchen.git
    # presention is located in ${GOPATH}/github.com/MarcGrol/goopenkitchen
    
*Step 3 - Start presentation within the goopenkitchen directory*

    cd ${GOPATH}/src/github.com/MarcGrol/goopenkitchen
    # run present tool in background
    ${GOPATH}/bin/present -http=127.0.0.1:3999 &

*Step 4 - Point your browser at:*
http://127.0.0.1:3999

*Step 5 - Build all examples in your local environment*

    cd ${GOPATH}/src/github.com/MarcGrol/goopenkitchen
    go install ./...
    # all examples should now be in ${GOPATH}/bin as executables

