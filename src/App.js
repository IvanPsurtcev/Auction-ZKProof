import './App.css'
import { useState, useEffect } from 'react'
import logo from './assets/images/logo.svg'
import metamask from './assets/images/metamask.svg'
import image1 from './assets/images/1.png'
import vessel from './assets/images/vessel.png'
import warm from './assets/images/warm.png'
import hat from './assets/images/hat.png'

function App() {
  const [modal1, setModal1] = useState(false)
  const [modal2, setModal2] = useState(false)
  const [modal3, setModal3] = useState(false)
  const [modal4, setModal4] = useState(false)
  const [modal5, setModal5] = useState(false)
  const [modal6, setModal6] = useState(false)
  const [amount, setAmount] = useState('')
  const [currentCard, setCurrentCard] = useState('')
  const [needConnection, setNeedConnection] = useState(true)
  const [currentAccount, setCurrentAccount] = useState('')

  useEffect(() => {
    if (typeof currentAccount !== 'undefined' && currentAccount.length > 0) {
      setNeedConnection(false)
      return
    }
    setNeedConnection(true)
  }, [currentAccount])

  const connectWallet = async (event) => {
    event.preventDefault()

    if (currentAccount) return

    if (typeof window.ethereum !== 'undefined') {
      try {
        const response = await window.ethereum.request({
          method: 'eth_requestAccounts',
        })

        setCurrentAccount(response[0])
      } catch (e) {
        setCurrentAccount(undefined)
      }
    }
  }

  const cardClickHandler = (name) => {
    setCurrentCard(name)
    setModal1(true)
  }

  // здесь отправка транзакции
  const sendTransaction = async () => {}

  return (
    <div className="wrapper">
      <header className="header">
        <div className="header__container _container">
          <img src={logo} />
          {needConnection ? (
            <div
              className="header__button"
              onClick={connectWallet}
            >
              <img src={metamask} />
              <p>Подключить кошелек</p>
            </div>
          ) : (
            <div
              className="header__button header__button_small-text"
              onClick={connectWallet}
            >
              <img src={metamask} />
              <p>{currentAccount}</p>
            </div>
          )}
        </div>
      </header>
      <main className="main">
        <div className="main__container _container">
          <div className="main__title">ZK блокчейн-аукцион</div>
          <div className="main__section section">
            <div className="section__description">
              <p>
                Электронный аукцион с закрытыми ставками <br />
                на блокчейне Ethereum со смарт-контрактом <br />и
                доказательством с нулевым разглашением
              </p>
            </div>
            <div className="section__image">
              <img src={image1} />
            </div>
          </div>
          <div className="main__cards-section cards-section">
            <div className="cards-section__title">Разыгрываем товары</div>
            <div className="cards-section__list">
              <div className="cards-section__item card">
                <div className="card__block">
                  <div className="card__image">
                    <img src={vessel} />
                  </div>
                  <div className="card__description">
                    Итальянская ваза Майолика, 15 век <br />
                    <b>Начальная ставка: 1 ETH Окончание торгов: 20.07.2023</b>
                  </div>
                </div>

                <div
                  className="card__button"
                  onClick={() => cardClickHandler(1)}
                >
                  сделать ставку
                </div>
              </div>

              <div className="cards-section__item card">
                <div className="card__block">
                  <div className="card__image">
                    <img src={warm} />
                  </div>
                  <div className="card__description">
                    NFT CryptoPunks
                    <br />
                    <b>
                      Начальная ставка: 0.01 ETH Окончание торгов: 23.05.2023
                    </b>
                  </div>
                </div>

                <div
                  className="card__button"
                  onClick={() => cardClickHandler(2)}
                >
                  сделать ставку
                </div>
              </div>

              <div className="cards-section__item card">
                <div className="card__block">
                  <div className="card__image">
                    <img src={hat} />
                  </div>
                  <div className="card__description">
                    Шапка Мономаха, 16 век <br />
                    <b>
                      Начальная ставка: 10 ETH Окончание торгов: 31.12.2023{' '}
                    </b>
                  </div>
                </div>

                <div
                  className="card__button"
                  onClick={() => cardClickHandler(3)}
                >
                  сделать ставку
                </div>
              </div>
            </div>
          </div>
        </div>
      </main>
      <footer className="footer">
        <div className="footer__container _container">
          <div className="footer__button">Начать аукцион</div>
          <div className="footer__button">Закончить аукцион</div>
          <div className="footer__button">Проверить выигрыш</div>
        </div>
      </footer>
      {modal1 && (
        <div className="modal">
          <div className="modal__body modal__body_grey">
            <div className="modal__header">
              <div
                className="modal__close"
                onClick={() => setModal1(false)}
              >
                &times;
              </div>
            </div>
            <div className="modal__main">
              <div>
                <div className="modal__input">
                  <input
                    placeholder="Введите сумму"
                    type="text"
                    value={amount}
                    onChange={(event) => setAmount(event.target.value)}
                  />
                  <p>ETH</p>
                </div>
                <div className="modal__info"></div>
              </div>
              <div className="modal__button">Сделать ставку</div>
            </div>
          </div>
        </div>
      )}
      {modal2 && (
        <div className="modal">
          <div className="modal__body modal__body_green">
            <div className="modal__header">
              <div
                className="modal__close"
                onClick={() => setModal2(false)}
              >
                &times;
              </div>
            </div>
            <div className="modal__text">
              Вы выиграли в этом аукционе, поздравляем!
            </div>
          </div>
        </div>
      )}
      {modal3 && (
        <div className="modal">
          <div className="modal__body modal__body_grey">
            <div className="modal__header">
              <div
                className="modal__close"
                onClick={() => setModal3(false)}
              >
                &times;
              </div>
            </div>
            <div className="modal__text">
              К сожалению, вы проиграли в этом аукционе.
            </div>
          </div>
        </div>
      )}
      {modal4 && (
        <div className="modal">
          <div className="modal__body modal__body_grey">
            <div className="modal__header">
              <div
                className="modal__close"
                onClick={() => setModal4(false)}
              >
                &times;
              </div>
            </div>
            <div className="modal__text">
              Торги еще не закончились, проверьте позже.
            </div>
          </div>
        </div>
      )}
      {modal5 && (
        <div className="modal">
          <div className="modal__body modal__body_grey">
            <div className="modal__header">
              <div
                className="modal__close"
                onClick={() => setModal5(false)}
              >
                &times;
              </div>
            </div>
            <div className="modal__text">
              Ваша ставка принята! <br />
              Ожидайте окончания торгов.
            </div>
          </div>
        </div>
      )}
      {modal6 && (
        <div className="modal">
          <div className="modal__body modal__body_grey">
            <div className="modal__header">
              <div
                className="modal__close"
                onClick={() => setModal6(false)}
              >
                &times;
              </div>
            </div>
          </div>
        </div>
      )}
    </div>
  )
}

export default App
