const fs = require("fs")

async function main(){
    const file = fs.createReadStream("big.csv")

    // const file = fs.readFileSync('small.csv', 'utf8')
    const requestInit = {
        body: file,
        keepalive: false,
        method: "POST",
    }

    const resp = await fetch("http://localhost:8080", requestInit)

    console.log(resp.body)
}

main().catch((err) => {
    console.error(err)
    process.exit(1)
})