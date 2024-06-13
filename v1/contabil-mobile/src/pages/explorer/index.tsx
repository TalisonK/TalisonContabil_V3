import React, { useState } from 'react';
import { View, Modal, StyleSheet, Button,Text, TextInput, TouchableOpacity } from 'react-native';
import Item from '../components/item';
import { filterActivity } from '../../services/activity';
import { ScrollView } from 'react-native-gesture-handler';
import { FontAwesomeIcon } from '@fortawesome/react-native-fontawesome';
import { faFilter, faMagnifyingGlass } from '@fortawesome/free-solid-svg-icons';
import { Activity } from '../../interfaces/Activity.interface';

const Explorer: React.FC = () => {

    const [resultado, setResultado] = useState<Activity[]>([])
    const [descricao, setDescricao] = useState("")
    const [metodo, setMetodo] = useState("")
    const [categoria, setCategoria] = useState("")
    const [user, setUser] = useState("")
    const [tipo, setTipo] = useState("")
    const [valor, setValor] = useState(0)
    const [dataInicio, setDataInicio] = useState(null)
    const [dataFim, setDataFim] = useState(null)

    const [modal, showModal] = useState(false)
    const [total, setTotal] = useState(0)

    const handleFilter = async () => {

        if(descricao.length < 3){
            return;
        }

        const data = {
            descricao: descricao.toLocaleLowerCase(),
            metodo,
            categoria,
            user,
            tipo,
            valor,
            dataInicio,
            dataFim,
        }

        const ret = await filterActivity(data);

        setResultado(ret.data as Array<never>)

        let soma = 0;
        for(let i of ret.data as Array<Activity>){
            soma += i.valor;
        }
        setTotal(soma);
    }

    return (
        <View style={styles.body}>
            <View style={styles.filterBay}>
                <TextInput style={styles.filterBayText} value={descricao} onChangeText={setDescricao} onSubmitEditing={handleFilter}/>
                <TouchableOpacity style={styles.filterGlass} onPress={handleFilter}><FontAwesomeIcon color='#FFF' icon={faMagnifyingGlass} /></TouchableOpacity>
                <TouchableOpacity style={styles.filterBayButton} onPress={() => {showModal(!modal)}}>
                    <FontAwesomeIcon icon={faFilter} size={23} style={styles.filterBayIcon} />
                </TouchableOpacity>
            </View>
            <ScrollView style={styles.listBody}>
                {resultado.map(e => <Item body={e}></Item>)}
            </ScrollView>
            <View style={styles.footer}>
                <Text style={styles.footerText}>Total: {total.toFixed(2)}</Text>
            </View>
            <Modal animationType='slide' transparent={false} visible={modal} onRequestClose={() => { showModal(false) }}>
                <View>
                    <Text>OI</Text>
                </View>
            </Modal>
        </View>
    )
}

const styles = StyleSheet.create({
    body:{
    },
    listBody: {
        paddingStart: 10,
        paddingEnd: 10,
        marginEnd: 10,
        height: "87%",
        width: "100%"
    },
    filterBay: {
        flexDirection: 'row',
        padding: 5,
        alignItems: 'center',
        height: 40
    },
    filterBayText: {
        width: "90%",
        height: 40,
        borderWidth: 1,
        borderRadius: 3,
        paddingStart:10
    },
    filterBayButton: {
        width: 40,
        height: 40,
        color: '#ffffff00',
        borderRadius: 10,
        justifyContent: 'center',
        alignContent: 'center'
    },
    filterBayIcon: {
        alignSelf: 'center'
    },
    filterGlass:{
        marginStart:-50,
        backgroundColor:"#7322a3",
        width:50,
        height:40,
        borderRadius:5,
        alignItems: 'center',
        justifyContent: 'center'
    },
    footer:{
        width:"100%",
        backgroundColor:'#9996',
        height:40,
        justifyContent:'center',
        marginTop:10
    },
    footerText:{
        fontWeight: 'bold',
        marginStart: 'auto',
        marginEnd: 10,
        fontSize: 20,
        textAlign: 'center'
    }
})

export default Explorer;