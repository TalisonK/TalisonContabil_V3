import {Injectable} from "@angular/core";


@Injectable()
export class MethodsUtil{
    getMethods(){
        return [
            "Cartão de Credito",
            "Cartão de Debito",
            "Dinheiro",
            "Transferencia",
            "PIX"
        ]
    }

    translateMethod(method: string){
        switch(method){
            case "Cartão de Credito":
                return "CREDIT_CARD";
            case "Cartão de Debito":
                return "DEBIT_CARD";
            case "Dinheiro":
                return "MONEY";
            case "Transferencia":
                return "TRANSFER";
            case "PIX":
                return "PIX";
        }
        return "";
    }

}
