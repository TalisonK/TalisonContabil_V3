

export const dataHandler = (dP: string):string => {
    let aux = ""
    const dataPagamento:Date = new Date(dP)
    const month = dataPagamento.getMonth() + 1;
    aux = dataPagamento.getDate() + "/" + (month < 10? ("0" + month): month) + "/" + dataPagamento.getFullYear();
    return aux;
}

export const monthHandler = (data: Date): string => {
    const aux = String(data.toString()).substring(4,7) + ", " + String(data.toString()).substring(10,16);
    return aux;
}

export const valueHandler = (valor:number):string => {

    valor = Math.round(valor * 100) / 100    

    let a = "";
    
    if(!String(a).includes(".")){
        a = String(valor)+ ".00";
    }
    else{
        a = String(valor)
    }

    let aux:any = String(a).split(".");

    let inicio = "";
    let fim = "";


    let cont = 0;

    for(let i = aux[0].length-1; i >= 0; i--){
        if(cont === 3){
            inicio = aux[0][i] + "." + inicio
            cont = 1;
        } else{
            inicio = aux[0][i] + inicio
            cont++;
        }
    }
    if(aux[1].length === 1){
        fim = aux[1] + "0"
    } else {
        fim = aux[1]
    }
    
    return inicio + "," + fim

}