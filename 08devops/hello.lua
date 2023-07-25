print("hello")

local Person = {
    name = "zhangsan",
    age = 14,
    hello = function(name, age)
        print(name, age)
    end
}

return Person
