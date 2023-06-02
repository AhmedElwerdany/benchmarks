const fs = require("fs")

const file = fs.readFileSync("../data.json", "utf-8")

const performance1 = performance.now()

const result = JSON.parse(file)

console.log(performance.now() - performance1)
