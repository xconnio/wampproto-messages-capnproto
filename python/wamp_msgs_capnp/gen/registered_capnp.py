"""This is an automatically generated stub for `registered.capnp`."""
import os

import capnp  # type: ignore

capnp.remove_import_hook()
here = os.path.dirname(os.path.abspath(__file__))
module_file = os.path.abspath(os.path.join(here, "registered.capnp"))
module = capnp.load(module_file)  # pylint: disable=no-member
Registered = module.Registered
RegisteredBuilder = Registered
RegisteredReader = Registered
