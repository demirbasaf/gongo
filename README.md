# gongo

Easy to use model-based Go ORM for MongoDB

``` 
gongo.Connect()

catModel := gongo.Model(domain.Cat{})
cats := catModel.Find()

dog := domain.Dog{ Name: "Max" }
dogModel := gongo.Model(domain.Dog{})
dogModel.New(dog)
```