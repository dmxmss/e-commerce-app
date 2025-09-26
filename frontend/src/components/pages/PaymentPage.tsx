import React from "react";
import { loadStripe } from "@stripe/stripe-js";
import { Elements, useStripe, useElements, PaymentElement } from "@stripe/react-stripe-js";
import { config } from "../../config.ts";
import { useLocation } from "react-router-dom";

const stripePromise = loadStripe(config.stripePublicKey);

const CheckoutForm = () => {
  const stripe = useStripe();
  const elements = useElements();

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();

    if (!stripe || !elements ) return;
    
    const { error } = await stripe.confirmPayment({
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
    <form className="p-15 flex flex-col" onSubmit={handleSubmit}>
      <PaymentElement />
      <button 
        className="m-a max-w-30 mt-10 px-4 py-1 bg-green-600 text-white rounded-lg hover:bg-green-700 transition"
        disabled={!stripe}
      >
        Buy
      </button>
    </form>
  );
}

const PaymentPage = () => {
  const location = useLocation();
  const clientSecret = location.state?.clientSecret;

  if (!clientSecret) {
    throw new Error("User did not provide client secret")
  }

  const options = { clientSecret };

  return (
    <Elements stripe={stripePromise} options={options}>
      <CheckoutForm />
    </Elements>
  )
}

export default PaymentPage;
