const fs = require("fs")
const fsPromise = require("fs/promises")

async function main(){
    // const filepath = "./example/big.csv"
    const filepath = "./example/small.csv"
    let { size } = await fsPromise.stat(filepath)

    const fileStream = fs.createReadStream(filepath)
    fileStream.once("end", ()=> fileStream.close())

    const requestInit = {
        body: fileStream,
        method: "POST",
        duplex: "half",
        // contentType: "text/plain",
    }

    let resp = await fetch("http://localhost:8080/stream-upload", requestInit)

    console.log(await resp.json())

    // file upload

    const form = new FormData()
    const buf = await fsPromise.readFile(filepath)
    const fileName = filepath.split("/")[2]

    form.append("file", new File(buf, fileName), fileName)

    requestInit.body = form
    resp = await fetch("http://localhost:8080/file-upload", requestInit)

    console.log(await resp.json())
    // console.log(form)
}

main().catch((err) => {
    console.error(err)
    process.exit(1)
})