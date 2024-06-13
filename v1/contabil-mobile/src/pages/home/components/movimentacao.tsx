import React, { useEffect, useState } from "react";
import { View, Text, StyleSheet, ScrollView } from "react-native";
import Item from "../../components/item";

const Movimentacao = (props: any) => {

    const [activities, setActivities] = useState([]);

    useEffect(()=>{
        if(props.activities){
            setActivities(props.activities)
        }
        
    },[props.activities])
    
    return (
        <ScrollView style={styles.container} alwaysBounceVertical={true} showsVerticalScrollIndicator={false}>
            <View >
                <Text style={styles.title}>Ultimas Movimentações</Text>
                {
                    activities.length > 0?
                        <View style={styles.body}>
                            {activities.map((e) => (<Item body={e} update={props.update}/>))}
                        </View>
                    :<View style={styles.NA}>
                        <Text >Nenhuma movimentação no mês</Text>
                    </View>
                }
                
            </View>
        </ScrollView>
    )
}

const styles = StyleSheet.create({
    container: {
        marginStart: 10,
        marginEnd: 10,
        marginBottom: 10,
        height: '70%',
        marginTop:3
    },
    title: {
        fontSize: 16,
        fontWeight: 'bold',
        marginBottom: 4,
        marginTop:15
    },
    body: {
        marginBottom: 30
    },
    NA:{
        justifyContent:'center',
        height:100,
        backgroundColor:'#dadada5e',
        textAlign:'center',
        alignItems: 'center',
        borderRadius:3
    }

})

export default Movimentacao;