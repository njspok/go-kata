# Domain Specific Language

Простая реализация DSL для описания солнечной системы.
Солнечная система состоит из планет с характеристиками, а вокруг планет могут быть спутники.

Пример: 

```golang
SolarSystem("Sun", func() {
    Name("MySun")
    Description("This is my home world.")

    Planet("Earth", func() {
        Description("This my home planet.")
        Mass(9999)

        Satellite("Moon", func() {
            Description("Beautiful thing!")
            Mass(111)
        })
    })

    Planet("Mars", func() {
        Description("This my feature planet.")
        Mass(8888)

        Satellite("Deimos", func() {
            Description("Rock")
            Mass(222)
        })

        Satellite("Phobos", func() {
            Description("Dead")
            Mass(121)
        })
    })
})
```