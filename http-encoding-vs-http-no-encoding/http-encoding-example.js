import https from "https";
import config from "./config.js";

const data = []

async function benchMe() {
    const time1 = performance.now();
    return new Promise((resolve, reject) => {
        https.get({
            hostname: new URL(config.html_url).hostname,
            headers: {
                "accept-encoding": "gzip, deflate, br"
            }
        }, (res) => {
            res.on("data", () => {})
            res.on('end', () => {
                resolve((performance.now() - time1))
            })
            res.on('error', reject)
        })
    })
}

for (let i = 0; i < 20; i++) {
 const time = await benchMe()
 data.push(time)
}

console.log(data.reduce((a, b) => a + b, 0) / data.length)