import React, {useState} from 'react';
import { View, StyleSheet } from 'react-native';
import Dropdown from '../components/dropdown';


const Chart:React.FC = () => {

    const [chartsList, _] = useState([
        {"nome":"teste1"},
        {"nome":"teste2"}
    ]);
    const [chart, setChart] = useState("");


    return(
        <View style={styles.container}>
            <Dropdown selected={chart} setSelection={setChart} filterSize={0} list={chartsList} listUpdater={async() =>{}} hasFilter={false}/>
        </View>
    )
}


const styles = StyleSheet.create({

    container:{
        marginTop:10,
        marginStart: 5,
        marginEnd: 5

    }

})

export default Chart