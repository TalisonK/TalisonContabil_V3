import React, { useEffect, useState } from 'react';
import { FlatList, SafeAreaView, StyleSheet, View, Text, TouchableOpacity, TextInput } from "react-native";
import { getFilteredCategory } from '../../services/category';
import { useAuth } from '../../contexts/auth';

interface Props {
    selected: string,
    setSelection(selectedValue: string): void,
    list: string[],
    listUpdater(selectedValue: string): Promise<void>,
    filterSize: number,
    hasFilter: boolean
}



const Dropdown = ({ selected, setSelection, list, listUpdater, filterSize, hasFilter }: Props) => {

    const [filteredData, setFilteredData] = useState([]);
    const [focus, OnFocus] = useState(false);

    const { verifyToken } = useAuth()

    useEffect(() => {
        verifyToken();
        if (selected.length >= filterSize - 1) {
            listUpdater(selected).then(() => { setFilteredData(list as never[]) });
        } else {
            setFilteredData([]);
        }
    }, [selected]);

    const ItemView = ({ item }: any) => (
        <TouchableOpacity activeOpacity={0.7} style={styles.textFilter} onPress={() => { setSelection(item.nome); OnFocus(false) }}>
            <Text >{item.nome}</Text>
        </TouchableOpacity>
    )

    const ItemSeparatorView = () => (<View style={styles.separator} />)

    return (
        <SafeAreaView>
            <View>
                <TextInput
                    onChangeText={setSelection}
                    value={selected}
                    style={styles.entries}
                    onFocus={() => { OnFocus(true) }}
                />
                    {
                    ((focus && selected.length >= filterSize && filteredData.length >= 1) || (!hasFilter && focus)) ?
                        <View style={styles.body}>
                            <FlatList
                                data={filteredData}
                                keyExtractor={(_, index) => index.toString()}
                                ItemSeparatorComponent={ItemSeparatorView}
                                renderItem={ItemView}
                            />
                        </View> :
                        <></>
                    }
            </View>
        </SafeAreaView>
    )
}

const styles = StyleSheet.create({
    separator: {
        height: 0.5,
        width: "100%",
        backgroundColor: '#c8c8c8'
    },
    textFilter: {
        padding: 5
    },
    entries: {
        fontSize: 16,
        borderWidth: 1,
        borderRadius: 5,
        paddingStart: 10,
        height: 50,
        backgroundColor: "#FFF"
    },
    body: {
        backgroundColor: "#FFF",
        width: "100%",
        borderWidth: 1,
        borderRadius: 3,
        paddingTop: 20,
        marginTop: -19,
        zIndex: -2
    }
})

export default Dropdown