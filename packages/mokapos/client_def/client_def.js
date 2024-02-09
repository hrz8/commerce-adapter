function main(params, ctx) {
    console.log(JSON.stringify(params))
    console.log("============")
    console.log(JSON.stringify(ctx))
    let name = args.name || 'stranger'
    let greeting = 'Hello Two ' + name + '!'
    console.log(greeting)
    return {"body": greeting}
  }

exports.main = main
