# v0.3.0

ENHANCEMENTS:

* resource/time_offset: Add `keepers` argument
* resource/time_rotating: Add `keepers` argument
* resource/time_static: Add `keepers` argument

# v0.2.0

BREAKING CHANGES:

* resource/time_static: The `expiration_` arguments have been moved to the new `time_rotating` resource as `rotation_` arguments.

FEATURES:

* **New Resource:** `time_offset`
* **New Resource:** `time_rotating`

# v0.1.0

FEATURES:

* **New Resource:** `time_static`