import React, { useState, useEffect } from "react";
import { View, StyleSheet, Text, Modal, Pressable, TouchableOpacity, Alert } from "react-native";
import { dataHandler, valueHandler } from "../../services/util";
import { FontAwesomeIcon } from "@fortawesome/react-native-fontawesome";
import { faCalendarDays, faPen, faX } from "@fortawesome/free-solid-svg-icons";
import { useAuth } from "../../contexts/auth";
import { adiantaActivity, deletaActivity } from "../../services/activity";

interface Props {
    body: {
        id:string,
        dataPagamento: string,
        descricao: string,
        categoriaName: string,
        parcelaAtual: number,
        parcelaTotal: number,
        tipo: string,
        valor: number,
        metodo: string,
        createdAt: string
    },
    update:VoidFunction
}

const Item = (props: Props) => {

    const { id, dataPagamento, descricao, categoriaName, parcelaAtual, parcelaTotal, tipo, valor, metodo, createdAt } = props.body

    const [modal, showModal] = useState(false)

    const [alertAdianta, setAlertAdianta] = useState(false)
    const [alertDelete, setAlertDelete] = useState(false)

    const { verifyToken } = useAuth()

    useEffect(() => {
        verifyToken();
        if(alertAdianta){
            setAlertAdianta(false);
            showModal(false);
            adiantaActivity(id).then(res => {
                setTimeout(() => {props.update(); console.log("vai carai");
                }, 2000);
                
            });
        }
    }, [alertAdianta])

    useEffect(() => {
        verifyToken();
        if(alertDelete){
            setAlertDelete(false);
            showModal(false);
            deletaActivity(id).then(res => {
                props.update();
            });
        }
    }, [alertDelete])

    const handlerAdianta = () => {

        Alert.alert("Confirmação do Adiantamento", "Tem certeza que deseja adiantar as proximas parcelas desta?", [
            {
                text: 'Cancelar',
                onPress: () => {setAlertAdianta(false)},
                style: 'cancel',
            },
            { 
                text: 'OK', 
                onPress: () => {setAlertAdianta(true)} 
            },
        ]);
        
    }

    const handlerDeleta = () => {
        Alert.alert("Confirmação da Remoção", "Tem certeza que deseja apagar essa atividade?", [
            {
                text: 'Cancelar',
                onPress: () => {setAlertDelete(false)},
                style: 'cancel',
            },
            { 
                text: 'OK', 
                onPress: () => {setAlertDelete(true)} 
            },
        ]);
    }

    return (
        <>
            <TouchableOpacity style={styles.container} onPress={() => { showModal(true) }}>
                <View style={{ width: 80 }}>
                    <Text style={styles.data}>{dataHandler(dataPagamento)}</Text>
                    <Text style={styles.descricao}>{descricao}</Text>
                    <Text style={styles.categoria}>{categoriaName}</Text>
                </View>
                <View>
                    {metodo === "Credito" ? <Text style={styles.parcelas}>{parcelaAtual} / {parcelaTotal}</Text> : <></>}
                </View>
                <View>
                    <Text style={tipo === "Income" ? styles.income : styles.expense}>{tipo === "Income" ? "" : "-"} R$ {valueHandler(valor)}</Text>
                </View>
            </TouchableOpacity>
            <Modal animationType='fade' transparent={true} visible={modal} onRequestClose={() => { showModal(false) }}>
                <View style={styles.modalContent}>
                    <View style={styles.modalHeader}>
                        <Text style={styles.modalTitulo}>{descricao}</Text>
                        <Pressable onPress={() => { showModal(!modal) }} style={styles.close}><FontAwesomeIcon icon={faX} size={22} /></Pressable>
                    </View>

                    <View style={styles.modalBody}>
                        <View style={styles.modalRow}>
                            <Text style={styles.strong}>dataPagamento: </Text>
                            <Text>{dataPagamento.substring(0, 10)}</Text>
                        </View>

                        <View style={styles.modalRow}>
                            <Text style={styles.strong}>tipo: </Text>
                            <Text>{tipo}</Text>
                        </View>

                        <View style={styles.modalRow}>
                            <Text style={styles.strong}>metodo: </Text>
                            <Text>{metodo}</Text>
                        </View>

                        {
                            metodo === "Credito" ?
                                <View style={styles.modalRow}>
                                    <Text style={styles.strong}>parcela: </Text>
                                    <Text>{parcelaAtual} de {parcelaTotal}</Text>
                                </View> :
                                <></>
                        }

                        <View style={styles.modalRow}>
                            <Text style={styles.strong}>categoria: </Text>
                            <Text>{categoriaName}</Text>
                        </View>

                        <View style={styles.modalRow}>
                            <Text style={styles.strong}>valor: </Text>
                            <Text>R$ {valueHandler(valor)}</Text>
                        </View>

                        <View style={styles.modalRow}>
                            <Text style={styles.strong}>createdAt: </Text>
                            <Text>{createdAt.substring(0, 19)}</Text>
                        </View>
                    </View>
                    <View style={styles.modalButtons}>
                        <TouchableOpacity style={{...styles.modalButtonsShape, backgroundColor: '#e7e42e'}} onPress={handlerAdianta}>
                            <FontAwesomeIcon icon={faCalendarDays} size={28}></FontAwesomeIcon>
                        </TouchableOpacity>
                        <TouchableOpacity style={{...styles.modalButtonsShape, backgroundColor: '#2ee72e'}}>
                            <FontAwesomeIcon icon={faPen} size={28}></FontAwesomeIcon>
                        </TouchableOpacity>
                        <TouchableOpacity style={{...styles.modalButtonsShape, backgroundColor: '#f02020'}} onPress={handlerDeleta}>
                            <FontAwesomeIcon icon={faX} size={28}></FontAwesomeIcon>
                        </TouchableOpacity>
                    </View>
                </View>
            </Modal>
        </>

    )
}


const styles = StyleSheet.create({
    container: {
        backgroundColor: 'rgba(240,240,240,0.8)',
        width: '100%',
        height: 70,
        flexDirection: 'row',
        justifyContent: 'space-between',
        alignItems: 'center'
    },
    data: {
        fontSize: 10,
        color: '#bdbcbc'
    },
    descricao: {
        fontWeight: 'bold',
        width: 140
    },
    categoria: {
        fontSize: 10,
    },
    income: {
        color: "#4de039"
    },
    expense: {
        color: "#f75a5a"
    },
    parcelas: {
        color: "#bdbcbc"
    },
    modalContent: {
        flex: 1,
        marginTop: 250,
        marginBottom: 180,
        backgroundColor: '#fffffffd',
        width: '70%',
        alignSelf: 'center',
        borderRadius: 10,
        paddingTop: 5,
        paddingBottom: 20,
        paddingStart: 5,
        paddingEnd: 5
    },
    modalTitulo: {
        fontWeight: 'bold',
        fontSize: 20,
        marginBottom: 20,
        paddingStart: 10
    },
    modalBody: {
        marginStart: 10,
        marginEnd: 10,
        marginBottom: 10
    },
    modalHeader: {
        flexDirection: 'row',
        alignContent: 'center',
        textAlign: 'center',
        justifyContent: 'space-between'
    },
    modalRow: {
        flexDirection: 'row',
        justifyContent: 'space-between'
    },
    close: {
        width: 30,
        height: 30,
        borderRadius: 30,
        alignItems: 'center',
        marginEnd: 2,
        justifyContent: 'center'
    },
    strong: {
        fontWeight: 'bold'
    },
    modalButtons:{
        flexDirection:'row',
        alignContent:'center',
        width:'100%',
        height:50,
        justifyContent:'space-around',
        marginTop: 5
    },
    modalButtonsShape:{
        paddingStart:15,
        paddingEnd: 15,
        paddingTop:30,
        paddingBottom:30,
        borderRadius:50,
        justifyContent:'center'
    }
})

export default Item;