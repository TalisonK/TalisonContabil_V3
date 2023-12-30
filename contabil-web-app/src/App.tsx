import React, { useEffect } from 'react';
import { RouterApp } from './router';
import { AppContainer } from './styles';
import Login from './pages/login/Login';
import { SnackbarProvider } from 'notistack';

function App() {

  const [user, setUser] = React.useState(null);


  useEffect(() => {
	const userStorage = localStorage.getItem("user");
	if(userStorage){
	  setUser(JSON.parse(userStorage));
	}
  }, [])

  return (

	<SnackbarProvider maxSnack={3} autoHideDuration={3000}>
		<AppContainer>
        	{user? <RouterApp/>: <Login/>}
    	</AppContainer>
	</SnackbarProvider>
    

  );
}

export default App;
