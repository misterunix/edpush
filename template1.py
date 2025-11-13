import sdkit
from sdkit.models import load_model
from sdkit.generate import generate_images
from sdkit.utils import log

context = sdkit.Context()
context.device: str = 'cpu'
context.num_inference_steps: int = 50
context.sampler: str = 'dpmpp_2s_a'

# set the path to the model file on the disk (.ckpt or .safetensors file)
context.model_paths['stable-diffusion'] = '/path/to/models/default-model'
load_model(context, 'stable-diffusion')

# generate the image

images = generate_images(context, prompt='A beautiful landscape', negative_prompt='low quality, blurry', seed=42, width=512 , height=512,num_inference_steps=50)

# save the image
images[0].save("small.png") # images is a list of PIL.Image

#log.info("Generated images!")