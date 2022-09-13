# gongo

A model-based Golang ORM for mongoDB

``` 
gongo.Connect()
userModel := gongo.Model(domain.User{})

users := userModel.Find()
```
