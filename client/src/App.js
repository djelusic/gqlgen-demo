import './App.scss';
import { gql, useQuery } from '@apollo/client';
import { useEffect } from 'react';
import { BrowserRouter as Router, Route, Switch, useHistory, useParams } from 'react-router-dom';

const OFFER_QUERY = gql`
  query offer {
    offer {
      events {
        id
        name
        markets {
          id
          name
          odds {
            id
            name
            value
          }
        }
      }
    }
  }
`

const OFFER_SUBSCRIPTION = gql`
  subscription {
    offer {
      events {
        id
        markets {
          id
          odds {
            id
            value
          }
        }
      }
    }
  }
`

function Offer() {
  const { data, loading, subscribeToMore } = useQuery(OFFER_QUERY);
  useEffect(() => {
    const sub = subscribeToMore({
      document: OFFER_SUBSCRIPTION,
    })
    return () => sub();
  }, [subscribeToMore])
  const history = useHistory();
  if (loading || !data) return null;
  return (
    <div className="offer">
      { data.offer.events.map(e => {
        return (
          <div className="market" key={e.id} onClick={() => history.push(`/event/${e.id}`)}>
            <div className="market-name">{e.name}</div>
            <div className="odds">
              { e.markets[0].odds.map(o => {
                return (
                  <div key={o.id} className="odd">
                    <div className="odd-name">{o.name}</div>
                    <div className="odd-value">{o.value.toFixed(2)}</div>
                  </div>
                );
              })}
            </div>
          </div>
        );
      })}
    </div>
  );
}

const EVENT_QUERY = gql`
  query event($id: Int!) {
    event(id: $id) {
      name
      markets {
        id
        name
        odds {
          id
          name
          value
        }
      }
    }
  }
`

const EVENT_SUBSCRIPTION = gql`
  subscription event($id: Int!) {
    event(id: $id) {
      id
      markets {
        id
        odds {
          id
          value
        }
      }
    }
  }
`

function Event() {
  const { id } = useParams();
  const { data, loading, subscribeToMore } = useQuery(
    EVENT_QUERY,
    {
      variables: {
        id,
      }
    }
  )
  useEffect(() => {
    const sub = subscribeToMore({
      document: EVENT_SUBSCRIPTION,
      variables: {
        id,
      }
    })
    return () => sub();
  }, [subscribeToMore, id])
  if (!data || loading) return null;
  const event = data.event;
  return (
    <div className="event">
      <div className="event-name">{event.name}</div>
      <div className="markets">
      { event.markets.map(market => {
        return (
          <div className="market" key={market.id}>
            <div className="market-name">{market.name}</div>
            <div className="odds">
              { market.odds.map(o => {
                return (
                  <div key={o.id} className="odd">
                    <div className="odd-name">{o.name}</div>
                    <div className="odd-value">{o.value.toFixed(2)}</div>
                  </div>
                );
              })}
            </div>
          </div>
        );
      })}
      </div>
    </div>
  )
}

function App() {
  return (
    <div className="App">
      <Router>
        <Switch>
          <Route path="/event/:id">
            <Event />
          </Route>
          <Route path="/">
            <Offer />
          </Route>
        </Switch>
      </Router>
    </div>
  );
}

export default App;
