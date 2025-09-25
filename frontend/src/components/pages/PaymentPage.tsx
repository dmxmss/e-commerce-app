import { loadStripe } from "@stripe/stripe-js";
import { Elements, useStripe, useElements, PaymentElement } from "@stripe/react-stripe-js";
import { config } from "../../config.ts";
import { useLocation } from "react-router-dom";

const stripePromise = loadStripe(config.stripePublicKey);

const CheckoutForm = () => {
  const stripe = useStripe();
  const elements = useElements();

  const handleSubmit = (e) => {
    e.preventDefault();

    if (!stripe || !elements ) return;
    
    const { error } = stripe.confirmPayment({
      elements,
      confirmParams: {
        return_url: config.host,
      }
    });

    if (error) {
      console.error(error.message);
    }
  };

  return (
    <form onSubmit={handleSubmit}>
      <PaymentElement />
      <button disabled={!stripe}>Buy</button>
    </form>
  );
}

const PaymentPage = () => {
  location = useLocation();
  const clientSecret = location.state?.clientSecret;

  if (!clientSecret) {
    throw new Error("User did not provide client secret")
  }

  const options = { clientSecret };

  return (
    <Elements stripe={stripePromise} options={options}>
      <CheckoutForm clientSecret={clientSecret} />
    </Elements>
  )
}

export default PaymentPage;
