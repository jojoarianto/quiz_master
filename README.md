# QUIZ MASTER

command line interface app for quiz

## how to play

run install/download depedencies  & unit testing

```bash
bin/setup
```
run shell app

```bash
bin/main
```

### project structure

```
    config/
        config.go
        constant.go
    domain/
        handler/
            questionhandler.go
        models/
            error.go
            question.go
        repository/
            questionrepo.go
        service
            questionsvc.go
    helper/
        string.go
        string_test.go
    infrastructure/
        sqlite3/
            questionrepo.go
            questionrepo_test.go
    interface/
        cli/
            handler/
                question.go
                question_test.go
                print.go
                print_test.go
            migration/
                main.go
    service/
        questionsvc.go
        questionsvc_test.go
    .gitignore
    main.go
    go.mod
    Makefile
```

### development mode

to run dev mode run
```bash
make run
```

to run test 
```bash
make test
```

see another make command on makefile
