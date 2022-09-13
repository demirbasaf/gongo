# gongo

Easy to use model-based Go ORM for MongoDB

``` 
gongo.Connect()
userModel := gongo.Model(domain.User{})

users := userModel.Find()
```
