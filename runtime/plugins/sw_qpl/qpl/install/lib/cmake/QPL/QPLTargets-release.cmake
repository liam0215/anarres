#----------------------------------------------------------------
# Generated CMake target import file for configuration "Release".
#----------------------------------------------------------------

# Commands may need to know the format version.
set(CMAKE_IMPORT_FILE_VERSION 1)

# Import target "QPL::qpl" for configuration "Release"
set_property(TARGET QPL::qpl APPEND PROPERTY IMPORTED_CONFIGURATIONS RELEASE)
set_target_properties(QPL::qpl PROPERTIES
  IMPORTED_LINK_INTERFACE_LANGUAGES_RELEASE "ASM_NASM;C;CXX"
  IMPORTED_LOCATION_RELEASE "${_IMPORT_PREFIX}/lib/libqpl.a"
  )

list(APPEND _cmake_import_check_targets QPL::qpl )
list(APPEND _cmake_import_check_files_for_QPL::qpl "${_IMPORT_PREFIX}/lib/libqpl.a" )

# Commands beyond this point should not need to know the version.
set(CMAKE_IMPORT_FILE_VERSION)
