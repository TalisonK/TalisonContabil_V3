import React, { useEffect, useState } from "react";
import { StyleSheet, View, TextInput, Text } from "react-native";
import DropDownPicker from 'react-native-dropdown-picker';
import { ScrollView } from "react-native-gesture-handler";
import { dataHandler } from "../../services/util";
import { DatePicker } from 'react-native-woodpicker'
import { FontAwesomeIcon } from "@fortawesome/react-native-fontawesome";
import { faCalendar } from "@fortawesome/free-solid-svg-icons";
import { Button } from "@react-native-material/core";
import { useAuth } from "../../contexts/auth";
import { sendActivity } from "../../services/activity";
import { showMessage } from "react-native-flash-message";
import SearchCategory from "../components/searchCategory";
import Dropdown from "../components/dropdown";
import { getFilteredCategory } from "../../services/category";

const InsertActivity: React.FC = ({ navigation, route }: any) => {

    const [descricao, setDescricao] = useState("");
    const [valor, setValor] = useState("");
    const [parcelaTotal, setPT] = useState(0);

    const [dataPagamento, setDP] = useState(new Date());
    const [showData, setShowData] = useState(false);

    const [tipo, setTipo] = useState("");
    const [showTipo, setShowTipo] = useState(false);

    const [metodo, setMetodo] = useState("");
    const [showMetodo, setShowMetodo] = useState(false);


    const [categoria, setCategoria] = useState("");
    const [filtroCategoria, setFiltroCategoria] = useState([]);
    const [showCategoria, setShowCategoria] = useState(false);

    useEffect(() => {
        setMetodo("");
    }, [tipo])

    useEffect(() => {
        setPT(0);
    }, [metodo])

    const [tipos, setTipos] = useState([
        { label: 'Income', value: 'Income' },
        { label: 'Expense', value: 'Expense' }
    ])

    const [metodos, setMetodos] = useState([
        { label: 'Credito', value: 'Credito' },
        { label: 'Debito', value: 'Debito' },
        { label: 'Dinheiro', value: 'Dinheiro' }
    ])

    const { user, verifyToken } = useAuth()

    const { update } = route.params;

    const handleSend = async () => {

        verifyToken();
        let auxDate: Date = new Date(dataPagamento);
        auxDate.setHours(3);

        const novo = {
            descricao: descricao.toLocaleLowerCase(),
            metodo: tipo === "Income" ? "Entrada" : metodo,
            categoriaName: tipo === "Income" ? "Income" : categoria,
            user,
            valor: Number(valor),
            tipo,


            dataPagamento: auxDate.toISOString(),
            parcelaAtual: metodo === "Credito" ? 1 : 0,
            parcelaTotal
        };

        const ret = await sendActivity(novo);

        if (ret.data?.status === true || ret.status == 200) {
            showMessage({
                message: "Movimento adicionado com sucesso",
                type: "success"
            })
        } else {

            showMessage({
                message: ret.data?.message !== undefined ? ret.data?.message : "Erro ao tentar adicionar",
                type: "danger"
            })
        }
        setTimeout(verifyToken, 1000);
        update();


    }

    const filterCategoria = async() => {
        const retorno = await getFilteredCategory(categoria);
        setFiltroCategoria(retorno as Array<never>);
    }

    return (
        <ScrollView style={styles.container} alwaysBounceVertical={true} showsVerticalScrollIndicator={false}>
            <View style={styles.container}>
                <View style={styles.box}>
                    <Text style={styles.text}>Descrição</Text>
                    <TextInput
                        onChangeText={setDescricao}
                        value={descricao}
                        style={styles.entries}
                    />
                </View>

                <View style={showTipo ? { ...styles.box, marginBottom: 80 } : styles.box}>
                    <Text style={styles.text}>Tipo</Text>
                    <DropDownPicker
                        open={showTipo}
                        value={tipo}
                        items={tipos}
                        setOpen={setShowTipo}
                        setValue={setTipo}
                        setItems={setTipos}
                        placeholder=""
                        style={styles.dropdown}
                    />
                </View>

                {
                    tipo === "Expense" ?
                        <>
                            <View style={showMetodo ? { ...styles.box, marginBottom: 120 } : styles.box}>
                                <Text style={styles.text}>Metodo</Text>
                                <DropDownPicker
                                    open={showMetodo}
                                    value={metodo}
                                    items={metodos}
                                    setOpen={setShowMetodo}
                                    setValue={setMetodo}
                                    setItems={setMetodos}
                                    style={styles.dropdown}
                                />
                            </View>
                            <View style={showCategoria ? { ...styles.box } : styles.box}>
                                <Text style={styles.text}>Categoria</Text>
                                <Dropdown selected={categoria} setSelection={setCategoria} list={filtroCategoria} filterSize={3} listUpdater={filterCategoria} hasFilter={true}/>
                            </View>
                        </>
                        : <></>
                }
                {
                    (tipo === "Expense" && metodo === "Credito") ?
                        <View style={styles.box}>
                            <Text style={styles.text}>N. Parcelas</Text>
                            <TextInput
                                onChangeText={(e) => { setPT(Number(e)) }}
                                value={String(parcelaTotal)}
                                style={styles.entries}
                                inputMode="numeric"
                            />
                        </View>
                        : <></>
                }
                <View style={styles.box}>
                    <Text style={styles.text}>Valor</Text>
                    <TextInput
                        onChangeText={setValor}
                        value={valor}
                        style={styles.entries}
                        inputMode="numeric"
                    />
                </View>
                <View>
                    <Text style={styles.text}>Data</Text>
                    <View style={styles.calendarRow}>
                        <TextInput
                            value={dataHandler(dataPagamento.toUTCString())}
                            style={{ ...styles.entries, width: "85%" }}
                            editable={false}
                        />
                        <DatePicker
                            style={styles.data}
                            onDateChange={(e: string) => { setDP(new Date(e)) }}
                            value={dataPagamento}
                            text={<FontAwesomeIcon color="#FFF" size={30} icon={faCalendar} />}
                            placeholder={"Data"}
                            onClose={() => { setShowData(false) }}
                        />
                    </View>
                </View>
                <Button style={styles.botaoEnv} title="Enviar" onPress={handleSend} />
            </View>
        </ScrollView>
    )
}


const styles = StyleSheet.create({
    container: {
        marginStart: 26,
        marginEnd: 26,
        marginBottom: 50
    },
    box: {
        marginTop: 15,
    },
    text: {
        marginStart: 15,
        marginTop: 2,
        margin: 5
    },
    entries: {
        fontSize: 16,
        borderWidth: 1,
        borderRadius: 5,
        paddingStart: 10,
        height: 50,
        backgroundColor: "#FFF"
    },
    dropdown: {
        height: 40,
        zIndex: 100
    },
    calendarRow: {
        flexDirection: 'row'
    },
    IconB: {
        width: 50,
        alignContent: 'center',
        justifyContent: 'center',
        marginStart: -10
    },
    icon: {
        width: 80,
    },
    data: {
        backgroundColor: "#653780",
        height: 50,
        width: 50,
        marginStart: -12,
        borderRadius: 5,
        paddingStart: 10
    },
    botaoEnv: {
        marginTop: 20,
        width: 100,
        alignSelf: 'center'
    }
})

export default InsertActivity;
