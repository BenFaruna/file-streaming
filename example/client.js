const fs = require("fs")

async function main(){
    const file = fs.createReadStream("big.csv")

    file.on("end", ()=> file.close())

    const requestInit = {
        body: file,
        method: "POST",
        duplex: "half",
    }

    const resp = await fetch("http://localhost:8080", requestInit)

    console.log(resp.body)
}

main().catch((err) => {
    console.error(err)
    process.exit(1)
})