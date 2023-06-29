// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: wso2/discovery/config/enforcer/jwt_generator.proto

package org.wso2.apk.enforcer.discovery.config.enforcer;

/**
 * <pre>
 * keypair model
 * </pre>
 *
 * Protobuf type {@code wso2.discovery.config.enforcer.Keypair}
 */
public final class Keypair extends
    com.google.protobuf.GeneratedMessageV3 implements
    // @@protoc_insertion_point(message_implements:wso2.discovery.config.enforcer.Keypair)
    KeypairOrBuilder {
private static final long serialVersionUID = 0L;
  // Use Keypair.newBuilder() to construct.
  private Keypair(com.google.protobuf.GeneratedMessageV3.Builder<?> builder) {
    super(builder);
  }
  private Keypair() {
    publicCertificatePath_ = "";
    privateKeyPath_ = "";
  }

  @java.lang.Override
  @SuppressWarnings({"unused"})
  protected java.lang.Object newInstance(
      UnusedPrivateParameter unused) {
    return new Keypair();
  }

  @java.lang.Override
  public final com.google.protobuf.UnknownFieldSet
  getUnknownFields() {
    return this.unknownFields;
  }
  private Keypair(
      com.google.protobuf.CodedInputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    this();
    if (extensionRegistry == null) {
      throw new java.lang.NullPointerException();
    }
    com.google.protobuf.UnknownFieldSet.Builder unknownFields =
        com.google.protobuf.UnknownFieldSet.newBuilder();
    try {
      boolean done = false;
      while (!done) {
        int tag = input.readTag();
        switch (tag) {
          case 0:
            done = true;
            break;
          case 10: {
            java.lang.String s = input.readStringRequireUtf8();

            publicCertificatePath_ = s;
            break;
          }
          case 18: {
            java.lang.String s = input.readStringRequireUtf8();

            privateKeyPath_ = s;
            break;
          }
          case 24: {

            useForSigning_ = input.readBool();
            break;
          }
          default: {
            if (!parseUnknownField(
                input, unknownFields, extensionRegistry, tag)) {
              done = true;
            }
            break;
          }
        }
      }
    } catch (com.google.protobuf.InvalidProtocolBufferException e) {
      throw e.setUnfinishedMessage(this);
    } catch (java.io.IOException e) {
      throw new com.google.protobuf.InvalidProtocolBufferException(
          e).setUnfinishedMessage(this);
    } finally {
      this.unknownFields = unknownFields.build();
      makeExtensionsImmutable();
    }
  }
  public static final com.google.protobuf.Descriptors.Descriptor
      getDescriptor() {
    return org.wso2.apk.enforcer.discovery.config.enforcer.JWTGeneratorProto.internal_static_wso2_discovery_config_enforcer_Keypair_descriptor;
  }

  @java.lang.Override
  protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internalGetFieldAccessorTable() {
    return org.wso2.apk.enforcer.discovery.config.enforcer.JWTGeneratorProto.internal_static_wso2_discovery_config_enforcer_Keypair_fieldAccessorTable
        .ensureFieldAccessorsInitialized(
            org.wso2.apk.enforcer.discovery.config.enforcer.Keypair.class, org.wso2.apk.enforcer.discovery.config.enforcer.Keypair.Builder.class);
  }

  public static final int PUBLIC_CERTIFICATE_PATH_FIELD_NUMBER = 1;
  private volatile java.lang.Object publicCertificatePath_;
  /**
   * <code>string public_certificate_path = 1;</code>
   * @return The publicCertificatePath.
   */
  @java.lang.Override
  public java.lang.String getPublicCertificatePath() {
    java.lang.Object ref = publicCertificatePath_;
    if (ref instanceof java.lang.String) {
      return (java.lang.String) ref;
    } else {
      com.google.protobuf.ByteString bs = 
          (com.google.protobuf.ByteString) ref;
      java.lang.String s = bs.toStringUtf8();
      publicCertificatePath_ = s;
      return s;
    }
  }
  /**
   * <code>string public_certificate_path = 1;</code>
   * @return The bytes for publicCertificatePath.
   */
  @java.lang.Override
  public com.google.protobuf.ByteString
      getPublicCertificatePathBytes() {
    java.lang.Object ref = publicCertificatePath_;
    if (ref instanceof java.lang.String) {
      com.google.protobuf.ByteString b = 
          com.google.protobuf.ByteString.copyFromUtf8(
              (java.lang.String) ref);
      publicCertificatePath_ = b;
      return b;
    } else {
      return (com.google.protobuf.ByteString) ref;
    }
  }

  public static final int PRIVATE_KEY_PATH_FIELD_NUMBER = 2;
  private volatile java.lang.Object privateKeyPath_;
  /**
   * <code>string private_key_path = 2;</code>
   * @return The privateKeyPath.
   */
  @java.lang.Override
  public java.lang.String getPrivateKeyPath() {
    java.lang.Object ref = privateKeyPath_;
    if (ref instanceof java.lang.String) {
      return (java.lang.String) ref;
    } else {
      com.google.protobuf.ByteString bs = 
          (com.google.protobuf.ByteString) ref;
      java.lang.String s = bs.toStringUtf8();
      privateKeyPath_ = s;
      return s;
    }
  }
  /**
   * <code>string private_key_path = 2;</code>
   * @return The bytes for privateKeyPath.
   */
  @java.lang.Override
  public com.google.protobuf.ByteString
      getPrivateKeyPathBytes() {
    java.lang.Object ref = privateKeyPath_;
    if (ref instanceof java.lang.String) {
      com.google.protobuf.ByteString b = 
          com.google.protobuf.ByteString.copyFromUtf8(
              (java.lang.String) ref);
      privateKeyPath_ = b;
      return b;
    } else {
      return (com.google.protobuf.ByteString) ref;
    }
  }

  public static final int USE_FOR_SIGNING_FIELD_NUMBER = 3;
  private boolean useForSigning_;
  /**
   * <code>bool use_for_signing = 3;</code>
   * @return The useForSigning.
   */
  @java.lang.Override
  public boolean getUseForSigning() {
    return useForSigning_;
  }

  private byte memoizedIsInitialized = -1;
  @java.lang.Override
  public final boolean isInitialized() {
    byte isInitialized = memoizedIsInitialized;
    if (isInitialized == 1) return true;
    if (isInitialized == 0) return false;

    memoizedIsInitialized = 1;
    return true;
  }

  @java.lang.Override
  public void writeTo(com.google.protobuf.CodedOutputStream output)
                      throws java.io.IOException {
    if (!getPublicCertificatePathBytes().isEmpty()) {
      com.google.protobuf.GeneratedMessageV3.writeString(output, 1, publicCertificatePath_);
    }
    if (!getPrivateKeyPathBytes().isEmpty()) {
      com.google.protobuf.GeneratedMessageV3.writeString(output, 2, privateKeyPath_);
    }
    if (useForSigning_ != false) {
      output.writeBool(3, useForSigning_);
    }
    unknownFields.writeTo(output);
  }

  @java.lang.Override
  public int getSerializedSize() {
    int size = memoizedSize;
    if (size != -1) return size;

    size = 0;
    if (!getPublicCertificatePathBytes().isEmpty()) {
      size += com.google.protobuf.GeneratedMessageV3.computeStringSize(1, publicCertificatePath_);
    }
    if (!getPrivateKeyPathBytes().isEmpty()) {
      size += com.google.protobuf.GeneratedMessageV3.computeStringSize(2, privateKeyPath_);
    }
    if (useForSigning_ != false) {
      size += com.google.protobuf.CodedOutputStream
        .computeBoolSize(3, useForSigning_);
    }
    size += unknownFields.getSerializedSize();
    memoizedSize = size;
    return size;
  }

  @java.lang.Override
  public boolean equals(final java.lang.Object obj) {
    if (obj == this) {
     return true;
    }
    if (!(obj instanceof org.wso2.apk.enforcer.discovery.config.enforcer.Keypair)) {
      return super.equals(obj);
    }
    org.wso2.apk.enforcer.discovery.config.enforcer.Keypair other = (org.wso2.apk.enforcer.discovery.config.enforcer.Keypair) obj;

    if (!getPublicCertificatePath()
        .equals(other.getPublicCertificatePath())) return false;
    if (!getPrivateKeyPath()
        .equals(other.getPrivateKeyPath())) return false;
    if (getUseForSigning()
        != other.getUseForSigning()) return false;
    if (!unknownFields.equals(other.unknownFields)) return false;
    return true;
  }

  @java.lang.Override
  public int hashCode() {
    if (memoizedHashCode != 0) {
      return memoizedHashCode;
    }
    int hash = 41;
    hash = (19 * hash) + getDescriptor().hashCode();
    hash = (37 * hash) + PUBLIC_CERTIFICATE_PATH_FIELD_NUMBER;
    hash = (53 * hash) + getPublicCertificatePath().hashCode();
    hash = (37 * hash) + PRIVATE_KEY_PATH_FIELD_NUMBER;
    hash = (53 * hash) + getPrivateKeyPath().hashCode();
    hash = (37 * hash) + USE_FOR_SIGNING_FIELD_NUMBER;
    hash = (53 * hash) + com.google.protobuf.Internal.hashBoolean(
        getUseForSigning());
    hash = (29 * hash) + unknownFields.hashCode();
    memoizedHashCode = hash;
    return hash;
  }

  public static org.wso2.apk.enforcer.discovery.config.enforcer.Keypair parseFrom(
      java.nio.ByteBuffer data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static org.wso2.apk.enforcer.discovery.config.enforcer.Keypair parseFrom(
      java.nio.ByteBuffer data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static org.wso2.apk.enforcer.discovery.config.enforcer.Keypair parseFrom(
      com.google.protobuf.ByteString data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static org.wso2.apk.enforcer.discovery.config.enforcer.Keypair parseFrom(
      com.google.protobuf.ByteString data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static org.wso2.apk.enforcer.discovery.config.enforcer.Keypair parseFrom(byte[] data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static org.wso2.apk.enforcer.discovery.config.enforcer.Keypair parseFrom(
      byte[] data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static org.wso2.apk.enforcer.discovery.config.enforcer.Keypair parseFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static org.wso2.apk.enforcer.discovery.config.enforcer.Keypair parseFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input, extensionRegistry);
  }
  public static org.wso2.apk.enforcer.discovery.config.enforcer.Keypair parseDelimitedFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input);
  }
  public static org.wso2.apk.enforcer.discovery.config.enforcer.Keypair parseDelimitedFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input, extensionRegistry);
  }
  public static org.wso2.apk.enforcer.discovery.config.enforcer.Keypair parseFrom(
      com.google.protobuf.CodedInputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static org.wso2.apk.enforcer.discovery.config.enforcer.Keypair parseFrom(
      com.google.protobuf.CodedInputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input, extensionRegistry);
  }

  @java.lang.Override
  public Builder newBuilderForType() { return newBuilder(); }
  public static Builder newBuilder() {
    return DEFAULT_INSTANCE.toBuilder();
  }
  public static Builder newBuilder(org.wso2.apk.enforcer.discovery.config.enforcer.Keypair prototype) {
    return DEFAULT_INSTANCE.toBuilder().mergeFrom(prototype);
  }
  @java.lang.Override
  public Builder toBuilder() {
    return this == DEFAULT_INSTANCE
        ? new Builder() : new Builder().mergeFrom(this);
  }

  @java.lang.Override
  protected Builder newBuilderForType(
      com.google.protobuf.GeneratedMessageV3.BuilderParent parent) {
    Builder builder = new Builder(parent);
    return builder;
  }
  /**
   * <pre>
   * keypair model
   * </pre>
   *
   * Protobuf type {@code wso2.discovery.config.enforcer.Keypair}
   */
  public static final class Builder extends
      com.google.protobuf.GeneratedMessageV3.Builder<Builder> implements
      // @@protoc_insertion_point(builder_implements:wso2.discovery.config.enforcer.Keypair)
      org.wso2.apk.enforcer.discovery.config.enforcer.KeypairOrBuilder {
    public static final com.google.protobuf.Descriptors.Descriptor
        getDescriptor() {
      return org.wso2.apk.enforcer.discovery.config.enforcer.JWTGeneratorProto.internal_static_wso2_discovery_config_enforcer_Keypair_descriptor;
    }

    @java.lang.Override
    protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
        internalGetFieldAccessorTable() {
      return org.wso2.apk.enforcer.discovery.config.enforcer.JWTGeneratorProto.internal_static_wso2_discovery_config_enforcer_Keypair_fieldAccessorTable
          .ensureFieldAccessorsInitialized(
              org.wso2.apk.enforcer.discovery.config.enforcer.Keypair.class, org.wso2.apk.enforcer.discovery.config.enforcer.Keypair.Builder.class);
    }

    // Construct using org.wso2.apk.enforcer.discovery.config.enforcer.Keypair.newBuilder()
    private Builder() {
      maybeForceBuilderInitialization();
    }

    private Builder(
        com.google.protobuf.GeneratedMessageV3.BuilderParent parent) {
      super(parent);
      maybeForceBuilderInitialization();
    }
    private void maybeForceBuilderInitialization() {
      if (com.google.protobuf.GeneratedMessageV3
              .alwaysUseFieldBuilders) {
      }
    }
    @java.lang.Override
    public Builder clear() {
      super.clear();
      publicCertificatePath_ = "";

      privateKeyPath_ = "";

      useForSigning_ = false;

      return this;
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.Descriptor
        getDescriptorForType() {
      return org.wso2.apk.enforcer.discovery.config.enforcer.JWTGeneratorProto.internal_static_wso2_discovery_config_enforcer_Keypair_descriptor;
    }

    @java.lang.Override
    public org.wso2.apk.enforcer.discovery.config.enforcer.Keypair getDefaultInstanceForType() {
      return org.wso2.apk.enforcer.discovery.config.enforcer.Keypair.getDefaultInstance();
    }

    @java.lang.Override
    public org.wso2.apk.enforcer.discovery.config.enforcer.Keypair build() {
      org.wso2.apk.enforcer.discovery.config.enforcer.Keypair result = buildPartial();
      if (!result.isInitialized()) {
        throw newUninitializedMessageException(result);
      }
      return result;
    }

    @java.lang.Override
    public org.wso2.apk.enforcer.discovery.config.enforcer.Keypair buildPartial() {
      org.wso2.apk.enforcer.discovery.config.enforcer.Keypair result = new org.wso2.apk.enforcer.discovery.config.enforcer.Keypair(this);
      result.publicCertificatePath_ = publicCertificatePath_;
      result.privateKeyPath_ = privateKeyPath_;
      result.useForSigning_ = useForSigning_;
      onBuilt();
      return result;
    }

    @java.lang.Override
    public Builder clone() {
      return super.clone();
    }
    @java.lang.Override
    public Builder setField(
        com.google.protobuf.Descriptors.FieldDescriptor field,
        java.lang.Object value) {
      return super.setField(field, value);
    }
    @java.lang.Override
    public Builder clearField(
        com.google.protobuf.Descriptors.FieldDescriptor field) {
      return super.clearField(field);
    }
    @java.lang.Override
    public Builder clearOneof(
        com.google.protobuf.Descriptors.OneofDescriptor oneof) {
      return super.clearOneof(oneof);
    }
    @java.lang.Override
    public Builder setRepeatedField(
        com.google.protobuf.Descriptors.FieldDescriptor field,
        int index, java.lang.Object value) {
      return super.setRepeatedField(field, index, value);
    }
    @java.lang.Override
    public Builder addRepeatedField(
        com.google.protobuf.Descriptors.FieldDescriptor field,
        java.lang.Object value) {
      return super.addRepeatedField(field, value);
    }
    @java.lang.Override
    public Builder mergeFrom(com.google.protobuf.Message other) {
      if (other instanceof org.wso2.apk.enforcer.discovery.config.enforcer.Keypair) {
        return mergeFrom((org.wso2.apk.enforcer.discovery.config.enforcer.Keypair)other);
      } else {
        super.mergeFrom(other);
        return this;
      }
    }

    public Builder mergeFrom(org.wso2.apk.enforcer.discovery.config.enforcer.Keypair other) {
      if (other == org.wso2.apk.enforcer.discovery.config.enforcer.Keypair.getDefaultInstance()) return this;
      if (!other.getPublicCertificatePath().isEmpty()) {
        publicCertificatePath_ = other.publicCertificatePath_;
        onChanged();
      }
      if (!other.getPrivateKeyPath().isEmpty()) {
        privateKeyPath_ = other.privateKeyPath_;
        onChanged();
      }
      if (other.getUseForSigning() != false) {
        setUseForSigning(other.getUseForSigning());
      }
      this.mergeUnknownFields(other.unknownFields);
      onChanged();
      return this;
    }

    @java.lang.Override
    public final boolean isInitialized() {
      return true;
    }

    @java.lang.Override
    public Builder mergeFrom(
        com.google.protobuf.CodedInputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws java.io.IOException {
      org.wso2.apk.enforcer.discovery.config.enforcer.Keypair parsedMessage = null;
      try {
        parsedMessage = PARSER.parsePartialFrom(input, extensionRegistry);
      } catch (com.google.protobuf.InvalidProtocolBufferException e) {
        parsedMessage = (org.wso2.apk.enforcer.discovery.config.enforcer.Keypair) e.getUnfinishedMessage();
        throw e.unwrapIOException();
      } finally {
        if (parsedMessage != null) {
          mergeFrom(parsedMessage);
        }
      }
      return this;
    }

    private java.lang.Object publicCertificatePath_ = "";
    /**
     * <code>string public_certificate_path = 1;</code>
     * @return The publicCertificatePath.
     */
    public java.lang.String getPublicCertificatePath() {
      java.lang.Object ref = publicCertificatePath_;
      if (!(ref instanceof java.lang.String)) {
        com.google.protobuf.ByteString bs =
            (com.google.protobuf.ByteString) ref;
        java.lang.String s = bs.toStringUtf8();
        publicCertificatePath_ = s;
        return s;
      } else {
        return (java.lang.String) ref;
      }
    }
    /**
     * <code>string public_certificate_path = 1;</code>
     * @return The bytes for publicCertificatePath.
     */
    public com.google.protobuf.ByteString
        getPublicCertificatePathBytes() {
      java.lang.Object ref = publicCertificatePath_;
      if (ref instanceof String) {
        com.google.protobuf.ByteString b = 
            com.google.protobuf.ByteString.copyFromUtf8(
                (java.lang.String) ref);
        publicCertificatePath_ = b;
        return b;
      } else {
        return (com.google.protobuf.ByteString) ref;
      }
    }
    /**
     * <code>string public_certificate_path = 1;</code>
     * @param value The publicCertificatePath to set.
     * @return This builder for chaining.
     */
    public Builder setPublicCertificatePath(
        java.lang.String value) {
      if (value == null) {
    throw new NullPointerException();
  }
  
      publicCertificatePath_ = value;
      onChanged();
      return this;
    }
    /**
     * <code>string public_certificate_path = 1;</code>
     * @return This builder for chaining.
     */
    public Builder clearPublicCertificatePath() {
      
      publicCertificatePath_ = getDefaultInstance().getPublicCertificatePath();
      onChanged();
      return this;
    }
    /**
     * <code>string public_certificate_path = 1;</code>
     * @param value The bytes for publicCertificatePath to set.
     * @return This builder for chaining.
     */
    public Builder setPublicCertificatePathBytes(
        com.google.protobuf.ByteString value) {
      if (value == null) {
    throw new NullPointerException();
  }
  checkByteStringIsUtf8(value);
      
      publicCertificatePath_ = value;
      onChanged();
      return this;
    }

    private java.lang.Object privateKeyPath_ = "";
    /**
     * <code>string private_key_path = 2;</code>
     * @return The privateKeyPath.
     */
    public java.lang.String getPrivateKeyPath() {
      java.lang.Object ref = privateKeyPath_;
      if (!(ref instanceof java.lang.String)) {
        com.google.protobuf.ByteString bs =
            (com.google.protobuf.ByteString) ref;
        java.lang.String s = bs.toStringUtf8();
        privateKeyPath_ = s;
        return s;
      } else {
        return (java.lang.String) ref;
      }
    }
    /**
     * <code>string private_key_path = 2;</code>
     * @return The bytes for privateKeyPath.
     */
    public com.google.protobuf.ByteString
        getPrivateKeyPathBytes() {
      java.lang.Object ref = privateKeyPath_;
      if (ref instanceof String) {
        com.google.protobuf.ByteString b = 
            com.google.protobuf.ByteString.copyFromUtf8(
                (java.lang.String) ref);
        privateKeyPath_ = b;
        return b;
      } else {
        return (com.google.protobuf.ByteString) ref;
      }
    }
    /**
     * <code>string private_key_path = 2;</code>
     * @param value The privateKeyPath to set.
     * @return This builder for chaining.
     */
    public Builder setPrivateKeyPath(
        java.lang.String value) {
      if (value == null) {
    throw new NullPointerException();
  }
  
      privateKeyPath_ = value;
      onChanged();
      return this;
    }
    /**
     * <code>string private_key_path = 2;</code>
     * @return This builder for chaining.
     */
    public Builder clearPrivateKeyPath() {
      
      privateKeyPath_ = getDefaultInstance().getPrivateKeyPath();
      onChanged();
      return this;
    }
    /**
     * <code>string private_key_path = 2;</code>
     * @param value The bytes for privateKeyPath to set.
     * @return This builder for chaining.
     */
    public Builder setPrivateKeyPathBytes(
        com.google.protobuf.ByteString value) {
      if (value == null) {
    throw new NullPointerException();
  }
  checkByteStringIsUtf8(value);
      
      privateKeyPath_ = value;
      onChanged();
      return this;
    }

    private boolean useForSigning_ ;
    /**
     * <code>bool use_for_signing = 3;</code>
     * @return The useForSigning.
     */
    @java.lang.Override
    public boolean getUseForSigning() {
      return useForSigning_;
    }
    /**
     * <code>bool use_for_signing = 3;</code>
     * @param value The useForSigning to set.
     * @return This builder for chaining.
     */
    public Builder setUseForSigning(boolean value) {
      
      useForSigning_ = value;
      onChanged();
      return this;
    }
    /**
     * <code>bool use_for_signing = 3;</code>
     * @return This builder for chaining.
     */
    public Builder clearUseForSigning() {
      
      useForSigning_ = false;
      onChanged();
      return this;
    }
    @java.lang.Override
    public final Builder setUnknownFields(
        final com.google.protobuf.UnknownFieldSet unknownFields) {
      return super.setUnknownFields(unknownFields);
    }

    @java.lang.Override
    public final Builder mergeUnknownFields(
        final com.google.protobuf.UnknownFieldSet unknownFields) {
      return super.mergeUnknownFields(unknownFields);
    }


    // @@protoc_insertion_point(builder_scope:wso2.discovery.config.enforcer.Keypair)
  }

  // @@protoc_insertion_point(class_scope:wso2.discovery.config.enforcer.Keypair)
  private static final org.wso2.apk.enforcer.discovery.config.enforcer.Keypair DEFAULT_INSTANCE;
  static {
    DEFAULT_INSTANCE = new org.wso2.apk.enforcer.discovery.config.enforcer.Keypair();
  }

  public static org.wso2.apk.enforcer.discovery.config.enforcer.Keypair getDefaultInstance() {
    return DEFAULT_INSTANCE;
  }

  private static final com.google.protobuf.Parser<Keypair>
      PARSER = new com.google.protobuf.AbstractParser<Keypair>() {
    @java.lang.Override
    public Keypair parsePartialFrom(
        com.google.protobuf.CodedInputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return new Keypair(input, extensionRegistry);
    }
  };

  public static com.google.protobuf.Parser<Keypair> parser() {
    return PARSER;
  }

  @java.lang.Override
  public com.google.protobuf.Parser<Keypair> getParserForType() {
    return PARSER;
  }

  @java.lang.Override
  public org.wso2.apk.enforcer.discovery.config.enforcer.Keypair getDefaultInstanceForType() {
    return DEFAULT_INSTANCE;
  }

}

