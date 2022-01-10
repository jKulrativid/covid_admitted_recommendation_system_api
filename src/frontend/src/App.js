import { Routes, Route } from 'react-router-dom'

import HomePage from './pages/Home'
import ResultPage from './pages/Result'

import Layout from './components/layout/Layout'

export default function App() {
  return <div>
    <Layout>
      <Routes>
        <Route exact path="/" element={HomePage} />
        <Route path="result" element={ResultPage} />
      </Routes>
    </Layout>
  </div>;
}
