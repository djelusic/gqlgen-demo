import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
import App from './App';
import reportWebVitals from './reportWebVitals';
import { ApolloClient, ApolloProvider, HttpLink, InMemoryCache, split } from '@apollo/client';
import { WebSocketLink } from '@apollo/client/link/ws';
import { getMainDefinition } from '@apollo/client/utilities';

const httpLink = new HttpLink({
  uri: 'http://localhost:8086/graphql',
});
const wsLink = new WebSocketLink({
  uri: 'ws://localhost:8086/graphql',
  options: {
    reconnect: true,
  },
});
const link = split(
  ({ query }) => {
    const { kind, operation } = getMainDefinition(query);
    return kind === 'OperationDefinition' && operation === 'subscription';
  },
  wsLink,
  httpLink,
);
const apolloClient = new ApolloClient({
  link: link,
  cache: new InMemoryCache({
    typePolicies: {
      Event: {
        merge: true,
        fields: {
          markets: {
            merge: mergeArrays("id"),
          },
        },
      },
      Market: {
        merge: true,
        fields: {
          odds: {
            merge: mergeArrays("id"),
          },
        },
      },
      Odd: {
        merge: true,
      },
    }
  }),
});

function mergeArrays(idField) {
  return (existing, incoming, { mergeObjects, readField }) => {
    const merged = existing ? existing.slice(0) : [];
    incoming.forEach(item => {
      const idx = merged.findIndex(el => readField(idField, el) === readField(idField, item))
      if (idx >= 0) {
        merged[idx] = mergeObjects(merged[idx], item);
      } else {
        merged.push(item);
      }
    })
    return merged;
  }
}

ReactDOM.render(
  <React.StrictMode>
    <ApolloProvider client={apolloClient}>
      <App />
    </ApolloProvider>
  </React.StrictMode>,
  document.getElementById('root')
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
