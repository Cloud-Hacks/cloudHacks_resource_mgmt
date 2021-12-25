const accountSid = os.environ['account_sid']; 
const authToken = os.environ['auth_token']; 
const to_whatsapp_no = os.environ['to_whatsapp_no'];
const from_whatsapp_no = os.environ['from_whatsapp_no'];

const client = require('twilio')(accountSid, authToken); 
 
client.messages 
      .create({ 
         body: 'Your Yummy Cupcakes Company order of 1 dozen frosted cupcakes has shipped and should be delivered today. Details: http://www.baristobakery.com/', 
         from: from_whatsapp_no,       
         to: to_whatsapp_no
 
       }) 
      .then(message => console.log(message.sid)) 
      .done();