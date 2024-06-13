import React from "react";
import { NavigationContainer } from '@react-navigation/native';

import Router from "./src/routes";
import { AuthProvider } from "./src/contexts/auth";
import FlashMessage from "react-native-flash-message";

const App: React.FC = () => {
  return (
    <>
      <NavigationContainer>
        <AuthProvider>
          <Router />
        </AuthProvider>
      </NavigationContainer>
      <FlashMessage position="bottom" />
    </>

  )
}

export default App;