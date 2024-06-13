import { Alert } from 'react-native';

const AlertDialog = (title: string, message: string):boolean => {

    let resultado = false;

    Alert.alert(title, message, [
        {
            text: 'Cancelar',
            onPress: () => {resultado = false},
            style: 'cancel',
        },
        { 
            text: 'OK', 
            onPress: () => {resultado = true} 
        },
    ]);

    return resultado;

}

export default AlertDialog;