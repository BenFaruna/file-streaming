const fs = require("fs")

async function main(){
    const file = fs.createReadStream("small.csv")
    file.once("end", ()=> file.close())

    const requestInit = {
        body: file,
        method: "POST",
        duplex: "half",
    }

    const resp = await fetch("http://localhost:8080/upload", requestInit)

    console.log(resp.body)
}

main().catch((err) => {
    console.error(err)
    process.exit(1)
})