const {prisma} = require("../../prisma/index")

//todos os user
export const allUsers = async() => {
    const allusers = await prisma.user.findMany()
    for (let u of allusers) delete u['senha']
    return allusers
}

//checar se user existe

export const checkUser =async (nome:string) => {
    const u = await prisma.user.findUnique({
        where:{nome}
    })
    return u ? true : false;
}

//criar user

export const createUser = async(nome:string, senha:string) => {

    const check = await checkUser(nome)
    
    if(!check){
        return await prisma.user.create({
            data:{
                nome,
                senha
            }
        })
    }else{
        return ""
    }
}

//alterar user

export const updateUser = async (user:string, nome:string, senha:string) => {

    const aux = await prisma.user.findMany({
        where:{nome}
    })
    for (let u of aux){
        if(u["id"] !== user){            
            return null
        }
    }
    const ret =  await prisma.user.update({
        where: {id: user},
        data:{
            nome,
            senha
        }
    })
    
    delete ret["senha"]
    return ret
}

// login

export const loginChecker = async (nome:string, senha:string) => {
    const aux = await prisma.user.findUnique({
        where:{
            nome
        }
    })

    if(aux === null){
        return null;
    }
    
    if(aux["senha"] === senha){
        return aux["id"]
    }
    else{
        return null;
    }

}