"""This is an automatically generated stub for `welcome.capnp`."""
import os

import capnp  # type: ignore

capnp.remove_import_hook()
here = os.path.dirname(os.path.abspath(__file__))
module_file = os.path.abspath(os.path.join(here, "welcome.capnp"))
module = capnp.load(module_file)  # pylint: disable=no-member
Welcome = module.Welcome
WelcomeBuilder = Welcome
WelcomeReader = Welcome
